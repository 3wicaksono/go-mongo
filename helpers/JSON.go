package helpers

import (
	"encoding/json"
)

// MakeJSON simple function to make JSON string
func MakeJSON(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(json)
}
