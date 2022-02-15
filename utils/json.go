package utils

import (
	"encoding/json"
	"fmt"
)

func InterfaceToJson(val interface{}) []byte {
	res, err := json.Marshal(val)
	if err != nil {
		fmt.Println("InterfaceToJson error: ", err)
		return nil
	}
	return res
}
