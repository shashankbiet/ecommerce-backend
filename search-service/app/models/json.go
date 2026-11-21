package model

import (
	"bytes"
	"encoding/json"
	"unicode"
)

func NormalizeJSON(data []byte) ([]byte, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	var value interface{}
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}
	return json.Marshal(normalizeJSONValue(value))
}

func normalizeJSONValue(value interface{}) interface{} {
	switch current := value.(type) {
	case map[string]interface{}:
		normalized := make(map[string]interface{}, len(current))
		for key, child := range current {
			normalizedKey := toSnakeCase(key)
			if normalizedKey != key {
				normalized[normalizedKey] = normalizeJSONValue(child)
			}
		}
		for key, child := range current {
			if toSnakeCase(key) == key {
				normalized[key] = normalizeJSONValue(child)
			}
		}
		return normalized
	case []interface{}:
		for index := range current {
			current[index] = normalizeJSONValue(current[index])
		}
		return current
	default:
		return value
	}
}

func toSnakeCase(value string) string {
	runes := []rune(value)
	result := make([]rune, 0, len(runes)+2)
	for index, current := range runes {
		if unicode.IsUpper(current) {
			previousIsLower := index > 0 && unicode.IsLower(runes[index-1])
			nextIsLower := index+1 < len(runes) && unicode.IsLower(runes[index+1])
			if index > 0 && (previousIsLower || nextIsLower) {
				result = append(result, '_')
			}
			current = unicode.ToLower(current)
		}
		result = append(result, current)
	}
	return string(result)
}
