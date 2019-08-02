package db

import (
  "context"
  "fmt"
  "log"
  "os"
  "time"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type store struct {
  db  *mongo.Database
}

var instance *store

func init() {
  // Set client options
  clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

  // Connect to MongoDB
  client, err := mongo.Connect(ctx, clientOptions)

  if err != nil {
    log.Fatal(err)
  }

  // Check the connection
  err = client.Ping(context.TODO(), nil)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Connected to MongoDB!")

  instance = &store{client.Database(os.Getenv("MONGODB_NAME"))}
}

func Close() error {
  return instance.db.Client().Disconnect(context.Background())
}
