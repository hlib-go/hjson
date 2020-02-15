package hjson

import (
	"encoding/json"
)

// Struct -> bytes
func StructToBytes(vstruct interface{}) []byte {
	bytes, err := json.Marshal(vstruct)
	if err != nil {
		panic(err)
	}
	return bytes
}

// Struct -> Json
func StructToJson(vstruct interface{}) string {
	return string(StructToBytes(vstruct))
}

// Struct -> Map
func StructToMap(vstruct interface{}, v interface{}) {
	JsonToMap(StructToJson(vstruct), v)
}
