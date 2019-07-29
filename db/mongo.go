package db

import (
  "context"
  "fmt"
  "log"
  "time"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
  *mongo.Database
}

func Connect(config *Config) (*Database, error) {
  // Set client options
  clientOptions := options.Client().ApplyURI(config.DatabaseURI)
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

  // Connect to MongoDB
  client, err := mongo.Connect(ctx, clientOptions)


  if err != nil {
   log.Fatal(err)
    return nil, err
  }

  // Check the connection
  err = client.Ping(context.TODO(), nil)

  if err != nil {
   log.Fatal(err)
   return nil, err
  }

  fmt.Println("Connected to MongoDB!")

  return &Database{client.Database(config.DatabaseName)}, nil
}
