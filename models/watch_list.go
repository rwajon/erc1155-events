package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchList struct {
	Id      primitive.ObjectID `json:"_id"`
	Address string             `json:"address" validate:"required"`
}
