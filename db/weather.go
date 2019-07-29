package db

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "time"

  "github.com/kenkwlai/weather-server/models"
  "go.mongodb.org/mongo-driver/bson"
)

func (db *Database) GetWeather(cityName string) (*models.CurrentWeatherData, error) {
  var weather models.CurrentWeatherData

  collection := db.Collection("weather")
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  filter := bson.M{
    "name": cityName,
  }
  findOptions := options.FindOne()
  findOptions.SetSort(bson.D{{"dt", -1}})

  err := collection.FindOne(ctx, filter, findOptions).Decode(&weather)

  if err != nil {
    return nil, err
  }

  return &weather, nil
}

func (db *Database) CreateWeather(weather *models.CurrentWeatherData) error {
  collection := db.Database.Collection("weather")
  ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
  res, err := collection.InsertOne(ctx, weather)

  if err != nil {
    log.Fatal(err)
    return err
  }

  id := res.InsertedID;
  log.Printf("Inserted with ID: %v", id)

  return nil
}
