package json

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Encode(data string) ([]string, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(data), &raw)
	if err != nil {
		return nil, errors.New("failed to unmarshal json string: " + err.Error())
	}

	var output []string
	for key, value := range raw {
		output = append(output, parse(value, key)...)
	}

	return output, nil
}

func parse(data interface{}, key string) []string {
	var output []string
	switch v := data.(type) {
	case string:
		output = append(output, fmt.Sprintf(`%s="%s"`, key, v))
	case int, int32, int64, float32, float64, bool:
		output = append(output, fmt.Sprintf("%s=%v", key, v))
	case []interface{}:
		for i, val := range v {
			output = append(output, parse(val, fmt.Sprintf("%s[%v]", key, i))...)
		}
	case map[string]interface{}:
		for k, val := range v {
			output = append(output, parse(val, fmt.Sprintf("%s/%s", key, k))...)
		}
	}

	return output
}
