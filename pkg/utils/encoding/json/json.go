package json

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

// JSONStruct json struct, same with runtime.RawExtension
type JSONStruct map[string]interface{}

// NewJSONStructByString new json struct from string
func NewJSONStructByString(source string) (*JSONStruct, error) {
	if source == "" {
		return nil, nil
	}
	var data JSONStruct
	err := json.Unmarshal([]byte(source), &data)
	if err != nil {
		return nil, fmt.Errorf("parse raw data failure %w", err)
	}
	return &data, nil
}

// NewJSONStructByStruct new json struct from struct object
func NewJSONStructByStruct(object interface{}) (*JSONStruct, error) {
	if object == nil {
		return nil, nil
	}
	var data JSONStruct
	out, err := yaml.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("marshal object data failure %w", err)
	}
	if err := yaml.Unmarshal(out, &data); err != nil {
		return nil, fmt.Errorf("unmarshal object data failure %w", err)
	}
	return &data, nil
}
