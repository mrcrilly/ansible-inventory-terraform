package main

import (
	"encoding/json"
	"errors"
	// "fmt"
	"strings"
)

import (
	"github.com/hashicorp/terraform/terraform"
)

type AWSProcessor struct {
	State *terraform.State

	inventory  map[string]*AnsibleInventoryGroup
	elasticIPs map[string]string
}

func (self *AWSProcessor) Process(t *terraform.State) error {
	if t == nil {
		return errors.New("Empty/invalid state provided.")
	}

	self.State = t
	self.elasticIPs = make(map[string]string, 0)

	if self.inventory == nil {
		self.inventory = make(map[string]*AnsibleInventoryGroup, 0)
	}

	for _, module := range t.Modules {
		for resourceKey, resourceValue := range module.Resources {
			switch resourceValue.Type {
			case "aws_instance":
				if len(module.Resources) == 0 {
					continue
				}

				groupName := module.Path[len(module.Path)-1]

				if _, ok := self.inventory[groupName]; !ok {
					self.inventory[groupName] = new(AnsibleInventoryGroup)
				}

				instanceName := strings.Split(resourceKey, ".")[1] + "-" + resourceValue.Primary.ID
				self.inventory[groupName].Hosts = append(self.inventory[groupName].Hosts, instanceName)
				break
			case "aws_eip":
				instanceName := strings.Split(resourceKey, ".")[1] + "-" + resourceValue.Primary.Attributes["instance"]
				self.elasticIPs[instanceName] = resourceValue.Primary.Attributes["public_ip"]
				break
			default:
				continue
			}
		}
	}

	return nil
}

func (self *AWSProcessor) Host(h string) (string, error) {
	var hostVariables map[string]interface{}

	if self.inventory == nil {
		return "", errors.New("Inventory has not been processed. Try Process().")
	}

	for _, module := range self.State.Modules {
		for resourceKey, resourceValue := range module.Resources {
			switch resourceValue.Type {
			case "aws_instance":
				instanceName := strings.Split(resourceKey, ".")[1] + "-" + resourceValue.Primary.ID

				if instanceName == h {
					hostVariables = make(map[string]interface{}, 0)

					if ipv4, ok := resourceValue.Primary.Attributes["public_ip"]; ok {
						if ipv4 != "" {
							hostVariables["ansible_ssh_host"] = ipv4
						} else {
							if eip, ok := self.elasticIPs[instanceName]; ok {
								hostVariables["ansible_ssh_host"] = eip
							} else {
								hostVariables["ansible_ssh_host"] = instanceName
							}
						}
					}

					if ipv4_private, ok := resourceValue.Primary.Attributes["private_ip"]; ok {
						hostVariables["private_ip"] = ipv4_private
					}
				}
			default:
				continue
			}
		}
	}

	if hostVariables == nil {
		return "{}", nil
	}

	j, err := json.Marshal(hostVariables)

	if err != nil {
		return "", err
	}

	return string(j), nil
}

func (self *AWSProcessor) Inventory() (string, error) {
	if self.inventory == nil {
		return "", errors.New("Inventory is empty. Use Process() to populate it.")
	}

	if len(self.inventory) == 0 {
		return "{}", nil
	}

	j, err := json.Marshal(self.inventory)

	if err != nil {
		return "", err
	}

	return string(j), nil
}

func (self *AWSProcessor) InventoryRaw() (map[string]*AnsibleInventoryGroup, error) {
	return nil, nil
}
