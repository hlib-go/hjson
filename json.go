package hjson

import (
	"encoding/json"
)

// Json 转 Map
func JsonToMap(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}

// Json 转 Struct
func JsonToStruct(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}
