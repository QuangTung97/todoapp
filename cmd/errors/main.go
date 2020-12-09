package main

import (
	"os"

	"gopkg.in/yaml.v2"
	"todoapp/lib/errors/generate"
)

func main() {
	file, err := os.Open("errors.yml")
	if err != nil {
		panic(err)
	}

	errorTags := make(map[string]generate.ErrorMap)

	err = yaml.NewDecoder(file).Decode(&errorTags)
	if err != nil {
		panic(err)
	}

	output, err := os.OpenFile("pkg/errors/errors.go", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	err = generate.Generate(errorTags, output)
	if err != nil {
		panic(err)
	}

	err = output.Close()
	if err != nil {
		panic(err)
	}
}
