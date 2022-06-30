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
		result := key
		switch v := value.(type) {
		case string:
			result += fmt.Sprintf(`="%s"`, v)
		case int, int32, int64, float32, float64:
			result += fmt.Sprintf("=%v", v)
		default:
			return nil, fmt.Errorf("can't handle a value of this type yet\n%v: %T", v, v)
		}
		output = append(output, result)
	}

	return output, nil
}
