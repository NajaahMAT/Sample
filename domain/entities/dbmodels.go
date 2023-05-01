package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SalesModel struct{
	ID     primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Sales  Sales           `json:"sales" bson:"sales"`
}



