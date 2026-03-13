package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	// omit bỏ qua
	// empty rỗng
	ID    bson.ObjectID `bson:"_id,omitempty"` // bson binary javascript object notation
	Name  string        `bson:"name"`          // json javascript object notation
	Email string        `bson:"email"`
}
