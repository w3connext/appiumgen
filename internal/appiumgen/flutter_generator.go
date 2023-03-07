package appiumgen

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type flutterGenerator struct {
}

// Generate implements Generator
func (f *flutterGenerator) Generate(content []byte) (string, error) {
	ymlMap := map[string]any{}
	err := yaml.Unmarshal(content, &ymlMap)
	if err != nil {
		fmt.Println("[ERROR] Invalid Syntax in", err)
	}

	fmt.Println(ymlMap)

	source := ""
	source = "Test"

	return source, nil
}

func NewFlutterGenerator() Generator {
	return &flutterGenerator{}
}
