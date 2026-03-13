package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func MongoDBConnect() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	client, _ := mongo.Connect(options.Client().ApplyURI(mongoURI))
	DB = client.Database(dbName)
}
