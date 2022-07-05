package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Decode(data []string) (string, error) {
	structuredData := make(map[string]interface{})
	dataList := parseStrings(data)
	fmt.Println("data list:", dataList)
	for _, v := range dataList {
		for key, val := range v {
			fmt.Println("key:", key)
			fmt.Println("val:", val)
			fmt.Println("structuredData[key]:", structuredData[key])
			if structuredData[key] == nil {
				fmt.Println("existing value not found")
				structuredData[key] = val
			} else if reflect.TypeOf(structuredData[key]).Kind() == reflect.Map && reflect.TypeOf(val).Kind() == reflect.Map {
				m := val.(map[string]interface{})
				for mKey, mVal := range m {
					structuredData[key].(map[string]interface{})[mKey] = mVal
				}
			} else {
				fmt.Printf("data exists at key %s: %v\n", key, structuredData[key])
			}
		}
	}

	bytes, err := json.Marshal(structuredData)
	if err != nil {
		return "", errors.New("failed to marshal json from map: " + err.Error())
	}
	return string(bytes), nil
}

func parseStrings(list []string) []map[string]interface{} {
	maps := make([]map[string]interface{}, 0)
	for _, val := range list {
		values := strings.Split(val, "=")
		k := values[0]

		v := values[1]
		value := strings.TrimPrefix(v, `"`)
		value = strings.TrimSuffix(value, `"`)

		maps = append(maps, parseString(k, value, maps...))
	}
	return maps
}

func parseString(key string, value interface{}, data ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	slashIndex := strings.Index(key, "/")
	leftBracketIndex := strings.Index(key, "[")
	if slashIndex > -1 {
		object := parseString(key[slashIndex+1:], value, data...)
		result[key[:slashIndex]] = object
		fmt.Println("result after slashIndex parsing:", result)
		// return result
	} else if leftBracketIndex > -1 {
		rightBracketIndex := strings.Index(key, "]")
		if rightBracketIndex > -1 {
			k := key[:leftBracketIndex]
			if rightBracketIndex+1 < len(key) {
				result[k] = parseString(key[rightBracketIndex+1:], value, data...)
			}
			found := false
			for _, v := range data {
				fmt.Println("v:", v)
				fmt.Println("k:", k)
				for mKey, mVal := range v {
					fmt.Println("mKey:", mKey)
					fmt.Println("mVal:", mVal)
					kind := reflect.TypeOf(mVal).Kind()
					if kind == reflect.Map {
						for mapKey, mapVal := range mVal.(map[string]interface{}) {
							if mapKey == k {
								d := make(map[string]interface{})
								d[mapKey] = mapVal
								result[mKey] = d
								found = true
							}
						}
					}
				}
				if v[k] != nil {
					kind := reflect.TypeOf(v[k]).Kind()
					if kind == reflect.Map {
						for mapKey, mapVal := range v[k].(map[string]interface{}) {
							fmt.Println("mapKey:", mapKey)
							fmt.Println("mapVal:", mapVal)
						}
					}
					if kind == reflect.Slice {
						fmt.Println("appending...")
						result[k] = append(v[k].([]interface{}), value)
					}
					found = true
				}
			}
			if !found {
				fmt.Println("replacing...")
				result[k] = []interface{}{value}
			}
			// return result
		}
	} else {
		result[key] = value
	}
	return result
}
