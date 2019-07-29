package api

import (
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/kenkwlai/weather-server/models"
  "log"
  "net/http"
)

const CityHk = "Hong Kong"
const Url = "https://api.openweathermap.org/data/2.5/weather?q=Hong Kong,hk&APPID=%s"

func (api *API) GetWeather(c *gin.Context) {
  res, err := http.Get(fmt.Sprintf(Url, api.Config.OpenWeatherApiKey))
  if err != nil || res.StatusCode != 200 {
    log.Println("Failed to get weather info from OpenWeatherMap")
    getFromDb(api, c)
    return
  }

  defer res.Body.Close()
  weatherData := new(models.CurrentWeatherData)
  err = json.NewDecoder(res.Body).Decode(&weatherData)
  if err != nil {
    log.Println("Failed to read weather info from OpenWeatherMap Api call")
    getFromDb(api, c)
    return
  }

  err = store(api, weatherData)
  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, weatherData)
}

func getFromDb(api *API, c *gin.Context) {
  weatherData, err := api.App.GetWeatherByCity(CityHk)
  if err != nil {
    c.Error(err)
  }

  c.JSON(http.StatusOK, weatherData)
}

func store(api *API, data *models.CurrentWeatherData) error {
  return api.App.CreateWeather(data)
}
