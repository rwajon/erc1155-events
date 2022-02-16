package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToInt(s string, params ...int) int64 {
	_params := [2]int{10, 64}
	for i, v := range params {
		_params[i] = v
	}

	if result, err := strconv.ParseInt(s, _params[0], _params[1]); err == nil {
		return result
	}
	return 0
}

func StringToFloat(s string, params ...int) float64 {
	_params := [2]int{64}
	for i, v := range params {
		_params[i] = v
	}

	if result, err := strconv.ParseFloat(s, _params[0]); err == nil {
		return result
	}
	return 0
}

func HexToInt(hex string) int64 {
	s := strings.Replace(hex, "0x", "", -1)
	s = strings.Replace(s, "0X", "", -1)
	return StringToInt(s, 16, 64)
}

func HexToFloat(hex string) float64 {
	s := fmt.Sprintf("%d", HexToInt(hex))
	return StringToFloat(s)
}
