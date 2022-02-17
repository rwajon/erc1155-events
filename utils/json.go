package utils

import (
	"encoding/json"
	"fmt"
)

func Jsonify(val interface{}) []byte {
	res, err := json.Marshal(val)
	if err != nil {
		fmt.Println("Jsonify error: ", err)
		return nil
	}
	return res
}
