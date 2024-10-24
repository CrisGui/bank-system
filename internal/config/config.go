package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "config/default.yaml"

func readConfigFile() map[interface{}]interface{} {
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func RunConfig() {
	fmt.Println("Hello from the config")
	for k, v := range readConfigFile() {
		fmt.Printf("%s -> %v\n", k, v)
	}
}
