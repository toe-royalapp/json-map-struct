package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

func StructToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	structValue := reflect.ValueOf(s).Elem()
	structType := structValue.Type()
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldName := structType.Field(i).Name
		result[fieldName] = field.Interface()
	}
	return result
}

func StructToJSON(s interface{}) (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func JsonToStruct(jsonData string, s interface{}) error {
	err := json.Unmarshal([]byte(jsonData), s)
	if err != nil {
		return err
	}
	return nil
}

func JsonToMap(jsonString string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MapToStruct(m map[string]interface{}, s interface{}) error {
	// Get the type of the struct
	structType := reflect.TypeOf(s).Elem()

	// Create a new struct value
	structValue := reflect.New(structType).Elem()

	// Loop through the map and set the struct fields
	for key, value := range m {
		fieldName := capitalizeFirstLetter(key)
		structField := structValue.FieldByName(fieldName)
		if structField.IsValid() && structField.CanSet() {
			fieldValue := reflect.ValueOf(value)
			structField.Set(fieldValue)
		}
	}

	// Set the struct pointer to the new struct value
	reflect.ValueOf(s).Elem().Set(structValue)

	return nil
}

// Helper function to capitalize the first letter of a string
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	first := strings.ToUpper(string(s[0]))
	if len(s) == 1 {
		return first
	}
	return first + s[1:]
}

func MapToJson(m map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
