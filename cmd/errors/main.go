package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type errorInfo struct {
	RPCStatus uint32            `yaml:"rpcStatus"`
	Code      string            `yaml:"code"`
	Message   string            `yaml:"message"`
	Details   map[string]string `yaml:"details"`
}

type errorList map[string]errorInfo

func main() {
	file, err := os.Open("errors.yml")
	if err != nil {
		panic(err)
	}

	errorTags := make(map[string]errorList)

	err = yaml.NewDecoder(file).Decode(&errorTags)
	if err != nil {
		panic(err)
	}

	fmt.Println(errorTags)
}
