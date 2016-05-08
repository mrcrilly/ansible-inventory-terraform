package main

import (
	"encoding/json"
	"fmt"
	"os"
)

import (
	"github.com/hashicorp/terraform/terraform"
)

type AnsibleInventory struct {
	Group     string
	Hosts     []string
	Variables map[string]interface{}
	Children  []string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	jsonRaw, err := os.Open("./terraform/terraform.tfstate")
	checkError(err)

	var terraformState *terraform.State

	jsonDecoder := json.NewDecoder(jsonRaw)
	err = jsonDecoder.Decode(&terraformState)
	checkError(err)

	var ansibleInventory AnsibleInventory = new(AnsibleInventory)

	for serverTFName, serverTFState := range terraformState.Modules[0].Resources {

	}
}
