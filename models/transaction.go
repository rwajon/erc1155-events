package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id              primitive.ObjectID `json:"_id"`
	BlockNumber     string             `json:"blockNumber"`
	BlockHash       string             `json:"blockHash" validate:"required"`
	Timestamp       int64              `json:"timestamp"`
	Date            time.Time          `json:"date"`
	From            string             `json:"from"`
	Gas             float64            `json:"gas"`
	GasPrice        float64            `json:"gasPrice"`
	Hash            string             `json:"hash" validate:"required"`
	To              string             `json:"to"`
	Type            uint8              `json:"type"`
	Value           float64            `json:"value"`
	SenderBalance   float64            `json:"senderBalance"`
	ReceiverBalance float64            `json:"receiverBalance"`
	ContractAddress string             `json:"contractAddress"`
}
