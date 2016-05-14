package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

import (
	"github.com/hashicorp/terraform/terraform"
)

type DigitalOceanProcessor struct {
	State     *terraform.State
	inventory map[string]*AnsibleInventoryGroup
}

func (self *DigitalOceanProcessor) Process(t *terraform.State) error {
	if t == nil {
		return errors.New("Empty/invalid state provided.")
	}

	self.State = t

	if self.inventory == nil {
		self.inventory = make(map[string]*AnsibleInventoryGroup, 0)
	}

	for _, module := range t.Modules {
		groupName := module.Path[len(module.Path)-1]

		if len(module.Resources) == 0 {
			continue
		}

		if _, ok := self.inventory[groupName]; !ok {
			self.inventory[groupName] = new(AnsibleInventoryGroup)
		}

		for _, resourceValue := range module.Resources {
			entry := fmt.Sprintf(
				"%s",
				resourceValue.Primary.Attributes["name"],
			)

			self.inventory[groupName].Hosts = append(self.inventory[groupName].Hosts, entry)
		}
	}

	return nil
}

func (self *DigitalOceanProcessor) Host(h string) (string, error) {
	var hostVariables map[string]interface{}

	if self.inventory == nil {
		return "", errors.New("Inventory has not been processed. Try Process().")
	}

	for _, module := range self.State.Modules {
		for _, resource := range module.Resources {
			if resource.Primary.Attributes["name"] == h {
				hostVariables = make(map[string]interface{}, 0)

				hostVariables["ansible_ssh_host"] = resource.Primary.Attributes["ipv4_address"]
				hostVariables["ansible_ssh_user"] = "root"
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

func (self *DigitalOceanProcessor) Inventory() (string, error) {
	if self.inventory == nil {
		return "", errors.New("Inventory is empty. Use Process() to populate it.")
	}

	if len(self.inventory) == 0 {
		// empty still requires a response to Ansible
		return "{}", nil
	}

	j, err := json.Marshal(self.inventory)

	if err != nil {
		return "", err
	}

	return string(j), nil
}

func (self *DigitalOceanProcessor) InventoryRaw() (map[string]*AnsibleInventoryGroup, error) {
	if self.inventory == nil || len(self.inventory) == 0 {
		return nil, errors.New("Inventory is empty. Use Process() to populate it.")
	}

	return self.inventory, nil
}
