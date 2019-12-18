package hjson

import (
	"encoding/json"
)

// Struct -> Json
func StructToJson(structData interface{}) string {
	bytes, err := json.Marshal(structData)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// Struct -> Map
func StructToMap(structData interface{}, v interface{}) {
	JsonToMap(StructToJson(structData), v)
}
