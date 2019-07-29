package app

import "github.com/kenkwlai/weather-server/models"

func (a *App) GetWeatherByCity(cityName string) (*models.CurrentWeatherData, error) {
  return a.Database.GetWeather(cityName)
}

func (a *App) CreateWeather(data *models.CurrentWeatherData) error {
  return a.Database.CreateWeather(data)
}
