package utils

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func YamlToJson(yamlInByte []byte) ([]byte, error) {
	var result interface{}
	if err := yaml.Unmarshal(yamlInByte, &result); err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

func JsonToYaml(jsonInByte []byte) ([]byte, error) {
	var result interface{}
	if err := json.Unmarshal(jsonInByte, &result); err != nil {
		return nil, err
	}
	return yaml.Marshal(result)
}
