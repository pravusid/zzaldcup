package helper

import (
	"encoding/json"
	"strconv"
)

func ParseInt(param string, defaultValue uint64) uint64 {
	str, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return defaultValue
	}
	return str
}

func ConvertJsonToMap(jsonString string) (map[string]interface{}, error) {
	var values map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &values)
	return values, err
}
