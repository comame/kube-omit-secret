package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var manifest SecretManifest
	err := yaml.Unmarshal([]byte(readStdin()), &manifest)
	if err != nil {
		panic(err)
	}

	for key := range manifest.Data {
		manifest.Data[key] = "~omit"
	}

	outputBytes, err := yaml.Marshal(manifest)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outputBytes))
}

type SecretManifest struct {
	ApiVersion string            `yaml:"apiVersion"`
	Data       map[string]string `yaml:"data"`
	Kind       string            `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Type string `yaml:"type"`
}

func readStdin() string {
	stdin := os.Stdin

	bytes, err := io.ReadAll(stdin)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func marshalYaml(v interface{}) string {
	bytes, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
