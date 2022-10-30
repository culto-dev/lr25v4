package utils

import (
	"encoding/json"
	"fmt"
)

func ObjectToJson(object any) []byte {
	fmt.Println(object)

	encoded, err := json.Marshal(object);
	if err != nil {
		fmt.Println("invalid object parsed")
		return []byte("[]")
	}
	return encoded
}