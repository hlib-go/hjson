package hjson

import (
	"encoding/json"
)

// Map -> bytes
func MapToBytes(vmap interface{}) []byte {
	bytes, err := json.Marshal(vmap)
	if err != nil {
		panic(err)
	}
	return bytes
}

// Map -> Json
func MapToJson(vmap interface{}) string {
	return string(MapToBytes(vmap))
}

// Map -> Struct
func MapToStruct(vmap interface{}, v interface{}) {
	JsonToStruct(MapToJson(vmap), v)
}
