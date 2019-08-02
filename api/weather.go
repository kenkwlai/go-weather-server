package api

import (
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/kenkwlai/weather-server/app"
  "github.com/kenkwlai/weather-server/models"
  "log"
  "net/http"
  "os"
)

const CityHk = "Hong Kong"
const Url = "https://api.openweathermap.org/data/2.5/weather?q=Hong Kong,hk&APPID=%s"

func getWeather(c *gin.Context) {
  res, err := http.Get(fmt.Sprintf(Url, os.Getenv("OPEN_WEATHER_MAP_API_KEY")))
  if err != nil || res.StatusCode != 200 {
    log.Println("Failed to get weather info from OpenWeatherMap")
    getFromDb(c)
    return
  }

  defer res.Body.Close()
  weatherData := new(models.CurrentWeatherData)
  err = json.NewDecoder(res.Body).Decode(&weatherData)
  if err != nil {
    log.Println("Failed to read weather info from OpenWeatherMap Api call")
    getFromDb(c)
    return
  }

  err = store(weatherData)
  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, weatherData)
}

func getFromDb(c *gin.Context) {
  weatherData, err := app.GetWeatherByCity(CityHk)
  if err != nil {
    c.Error(err)
    c.String(http.StatusInternalServerError, "internal server error")
    return
  }

  c.JSON(http.StatusOK, weatherData)
}

func store(data *models.CurrentWeatherData) error {
  return app.CreateWeather(data)
}
