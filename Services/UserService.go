package services

import (
	configs "SliceServer/Configs"
	models "SliceServer/Models"
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Users []models.User

const filePath = "Datas/UserData.json"

func GetCollection() *mongo.Collection {
	return configs.DB.Collection(os.Getenv("USER_COLLECTION"))
}

func GetAllUsers() []models.User {
	var users []models.User
	cursor, _ := GetCollection().Find(context.TODO(), bson.M{})
	cursor.All(context.TODO(), &users)
	return users
}

func Create(newUser models.User) models.User {
	newUser.ID = bson.NewObjectID()
	GetCollection().InsertOne(context.TODO(), newUser)
	return newUser
}

func GetUserById(id string) models.User {
	var user models.User
	objId, _ := bson.ObjectIDFromHex(id)
	GetCollection().FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&user)
	return user
}

func UserExistsById(id string) bool {
	user := GetUserById(id)
	if user.ID.IsZero() {
		return false
	}
	return true
}

func UpdateUserById(id string, updateUser models.User) models.User {
	var user models.User
	objId, _ := bson.ObjectIDFromHex(id)
	update := bson.M{
		"$set": bson.M{
			"name":  updateUser.Name,
			"email": updateUser.Email,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	GetCollection().FindOneAndUpdate(context.TODO(), bson.M{"_id": objId}, update, opts).Decode(&user)
	return user
}

func DeleteUser(id string) {
	objId, _ := bson.ObjectIDFromHex(id)
	GetCollection().FindOneAndDelete(context.TODO(), bson.M{"_id": objId})
}

// func LoadUsers() {
// 	file, _ := os.Open(filePath)
// 	defer file.Close()
// 	json.NewDecoder(file).Decode(&Users)
// }
// func SaveUsers() error {
// 	file, _ := os.Create(filePath)
// 	defer file.Close()
// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ")
// 	return encoder.Encode(Users)
// }
