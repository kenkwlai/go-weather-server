package db

import (
  "context"
  "fmt"
  "github.com/kenkwlai/weather-server/config"
  "log"
  "time"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type store struct {
  db  *mongo.Database
}

var instance *store

var mongoDbUri = config.GetOrDefault("MONGODB_URI", "mongodb://datastore:27017/default")
var mongoDbName = config.GetOrDefault("MONBODB_NAME", "default")

func Init() {
  // Set client options
  clientOptions := options.Client().ApplyURI(mongoDbUri)
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

  // Connect to MongoDB
  client, err := mongo.Connect(ctx, clientOptions)

  if err != nil {
    // Cannot connect, abort
    log.Fatal(err)
  }

  // Check the connection
  err = client.Ping(context.TODO(), nil)

  if err != nil {
    // Cannot connect, abort
    log.Fatal(err)
  }

  fmt.Println("Connected to MongoDB!")

  instance = &store{client.Database(mongoDbName)}
}

func Close() error {
  return instance.db.Client().Disconnect(context.Background())
}
