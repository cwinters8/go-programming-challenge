package json

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

func Decode(data []string) (string, error) {
	structuredData := make(map[string]interface{})
	for _, val := range data {
		values := strings.Split(val, "=")
		k := values[0]
		v := values[1]
		var key string
		index := 0
		leftBracketIndex := -1
		rightBracketIndex := -1
		for index < len(k) {
			char := string(k[index])
			if char == "[" {
				leftBracketIndex = index
			} else if char == "]" {
				rightBracketIndex = index
			}
			if leftBracketIndex > -1 && rightBracketIndex > -1 {
				break
			}
			index++
			if leftBracketIndex > -1 {
				continue
			}
			key += char
		}
		value := strings.TrimPrefix(v, `"`)
		value = strings.TrimSuffix(value, `"`)
		if leftBracketIndex > -1 && rightBracketIndex > -1 {
			if structuredData[key] != nil && reflect.TypeOf(structuredData[key]).Kind() == reflect.Slice {
				currentList := structuredData[key].([]interface{})
				structuredData[key] = append(currentList, value)
			} else {
				structuredData[key] = []interface{}{value}
			}
		} else {
			structuredData[key] = value
		}
	}

	bytes, err := json.Marshal(structuredData)
	if err != nil {
		return "", errors.New("failed to marshal json from map: " + err.Error())
	}
	return string(bytes), nil
}

// func parseKeyValue(key string, value string) interface{} {
// 	leftBracketIdx := -1
// 	rightBracketIdx := -1
// 	slashIdx := -1
// 	for i, v := range key {
// 		char := string(v)
// 		if leftBracketIdx == -1 && char == "[" {
// 			leftBracketIdx = i
// 		} else if rightBracketIdx == -1 && char == "]" {
// 			rightBracketIdx = i
// 		} else if slashIdx == -1 && char == "/" {
// 			slashIdx = i
// 		}
// 		if leftBracketIdx > -1 && rightBracketIdx > -1 && slashIdx > -1 {
// 			break
// 		}
// 	}
// 	if slashIdx > -1 {
// 		data := make(map[string]interface{})
// 		k := key[0:slashIdx]
// 		v := parseKeyValue(key[slashIdx+1:], value)
// 		data[k] = v
// 		return data
// 	}
// 	if leftBracketIdx > -1 && rightBracketIdx > -1 {
// 		data := []string{}
// 		k := key[slashIdx+1 : leftBracketIdx]
// 		index := int(key[leftBracketIdx+1])
// 		v := []string{value}

// 	}

// 	return nil
// }
