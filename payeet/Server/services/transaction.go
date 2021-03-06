package services

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Transaction struct contains info about a transaction.
type Transaction struct {
	Sender   string `bson:"Sender" json:"Sender"`
	Receiver string `bson:"Receiver" json:"Receiver"`
	Amount   int    `bson:"Amount" json:"Amount"`
	Time     int64  `bson:"Time" json:"Time"`
}

// ToBson truns a Transaction object into bson
func (Transaction *Transaction) ToBson() bson.D {

	a := bson.D{
		{Key: "Sender", Value: Transaction.Sender},
		{Key: "Receiver", Value: Transaction.Receiver},
		{Key: "Amount", Value: Transaction.Amount},
		{Key: "Time", Value: Transaction.Time},
	}

	return a
}
