package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var databaseName string

// -------------------- CONNECT DATABASE --------------------

func ConnectDB() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	mongoURI := os.Getenv("MONGO_URI")
	databaseName = os.Getenv("DB_NAME")

	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	DB = client

	log.Println("✅ MongoDB Connected Successfully")
}

// -------------------- GET COLLECTION --------------------

func GetCollection(collectionName string) *mongo.Collection {

	if DB == nil {
		log.Fatal("Database not initialized. Call ConnectDB() first.")
	}

	return DB.Database(databaseName).Collection(collectionName)
}

// -------------------- DISCONNECT DATABASE --------------------

func DisconnectDB() {

	if DB == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := DB.Disconnect(ctx); err != nil {
		log.Println("Error disconnecting MongoDB:", err)
	} else {
		log.Println("MongoDB disconnected")
	}
}
