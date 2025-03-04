package configfacets

import (
	"errors"
	"strings"
)

// getByPath retrieves a nested value from a JSON object using a dot-separated key path
func getByPath(data map[string]interface{}, path string) (interface{}, error) {
	keys := strings.Split(path, ".")
	var value interface{} = data

	for _, key := range keys {
		if m, ok := value.(map[string]interface{}); ok {
			if val, exists := m[key]; exists {
				value = val
			} else {
				return nil, errors.New("key not found: " + key)
			}
		} else {
			return nil, errors.New("invalid key path")
		}
	}

	return value, nil
}
