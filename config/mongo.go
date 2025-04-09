package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")

	if uri == "" || dbName == "" {
		log.Fatal("❌ Las variables MONGODB_URI o MONGODB_NAME no están definidas")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ Error creando el cliente de Mongo:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Error conectando a MongoDB:", err)
	}

	DB = client.Database(dbName)
	fmt.Println("✅ Conectado correctamente a MongoDB:", dbName)
}

func GetCollection(nombre string) *mongo.Collection {
	return DB.Collection(nombre)
}
