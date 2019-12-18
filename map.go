package hjson

import (
	"encoding/json"
)

// Map -> Json
func MapToJson(mapData interface{}) string {
	bytes, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// Map -> Struct
func MapToStruct(mapData interface{}, v interface{}) {
	JsonToStruct(MapToJson(mapData), v)
}
