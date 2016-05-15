package main

import (
	"github.com/hashicorp/terraform/terraform"
)

type AnsibleInventoryGroup struct {
	Hosts []string `json:"hosts"`
}

type AnsibleInventoryHost struct {
	Variables map[string]interface{}
}

type StateProcessor interface {
	Process(*terraform.State) error
	Host(string) (string, error)
	Inventory() (string, error)
	InventoryRaw() (map[string]*AnsibleInventoryGroup, error)
}
