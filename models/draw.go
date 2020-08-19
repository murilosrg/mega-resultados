package models

import (
	"time"
)

type Draw struct {
	Concourse           int       `json:"concourse" bson:"concourse"`
	Date                time.Time `json:"date" bson:"date"`
	Numbers             []int     `json:"numbers" bson:"numbers"`
	Collection          float64   `json:"collection" bson:"collection"`
	Accumulated         bool      `json:"accumulated" bson:"accumulated"`
	AccumulatedValue    float64   `json:"accumulatedValue" bson:"accumulatedValue"`
	EstimatedPrize      float64   `json:"estimatedPrize" bson:"estimatedPrize"`
	AccumulatedLastDraw float64   `json:"accumulatedLastDraw" bson:"accumulatedLastDraw"`
	Winners             *[]Winner `json:"winners" bson:"winners"`
}

type Winner struct {
	Number int     `json:"number" bson:"number"`
	Count  int     `json:"count" bson:"count"`
	Amount float64 `json:"amount" bson:"amount"`
}
