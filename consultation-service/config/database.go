package config

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	db := os.Getenv("MONGO_DB")
	if db == "" {
		log.Fatal("MONGO_DB environment variable is not set")
	}
	collection := client.Database(db).Collection(collectionName)
	return collection
}

func MySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/teman_sehat?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
