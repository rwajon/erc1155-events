package models

import "time"

type Transaction struct {
	BlockNumber     string    `json:"blockNumber"`
	BlockHash       string    `json:"blockHash"`
	Timestamp       int64     `json:"timestamp"`
	Date            time.Time `json:"date"`
	From            string    `json:"from"`
	Gas             float64   `json:"gas"`
	GasPrice        float64   `json:"gasPrice"`
	Hash            string    `json:"hash"`
	To              string    `json:"to"`
	Type            uint8     `json:"type"`
	Value           float64   `json:"value"`
	SenderBalance   float64   `json:"senderBalance"`
	ReceiverBalance float64   `json:"receiverBalance"`
	ContractAddress string    `json:"contractAddress"`
}
