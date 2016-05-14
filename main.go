package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

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
	Group(string) (string, error)
	Host(string) (string, error)
	Inventory() (string, error)
	InventoryRaw() (map[string]*AnsibleInventoryGroup, error)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	envStateFile := os.Getenv("ANSIBLE_INVENTORY_TERRAFORM_FILE")
	envProvider := os.Getenv("ANSIBLE_INVENTORY_TERRAFORM_PROVIDER")

	if envProvider == "" {
		envProvider = "digitalocean"
	}

	if envStateFile == "" {
		envStateFile = "./terraform.tfstate"
	}

	jsonRaw, err := os.Open(envStateFile)
	checkError(err)

	var terraformState *terraform.State

	jsonDecoder := json.NewDecoder(jsonRaw)
	err = jsonDecoder.Decode(&terraformState)
	checkError(err)

	var flagEverything = flag.Bool("list", true, "--list: will give you the entire inventory")
	var flagGroup = flag.String("group", "", "--group <group>: will give you a list of hosts in a group")
	var flagHost = flag.String("host", "", "--host <host>: will give host specific variables")

	flag.Parse()

	var processor StateProcessor
	var processorOut string

	switch envProvider {
	case "digitalocean":
		processor = new(DigitalOceanProcessor)
	default:
		panic(errors.New("No provider specified in ANSIBLE_INVENTORY_TERRAFORM_PROVIDER"))
	}

	processor.Process(terraformState)

	if *flagHost != "" {
		*flagEverything = false
		processorOut, err = processor.Host(*flagHost)
	}

	if *flagGroup != "" {
		*flagEverything = false
		processorOut, err = processor.Group(*flagGroup)
	}

	if *flagEverything {
		processorOut, err = processor.Inventory()
	}

	checkError(err)

	fmt.Printf("%s", processorOut)
	os.Exit(0)
}
