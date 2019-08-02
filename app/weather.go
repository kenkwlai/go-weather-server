package app

import (
  "github.com/kenkwlai/weather-server/db"
  "github.com/kenkwlai/weather-server/models"
)

func GetWeatherByCity(cityName string) (*models.CurrentWeatherData, error) {
  return db.WeatherMongoStore().GetWeather(cityName)
}

func CreateWeather(data *models.CurrentWeatherData) error {
  return db.WeatherMongoStore().CreateWeather(data)
}
