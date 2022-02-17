package models

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data"`
	Meta    map[string]interface{} `json:"meta"`
}
