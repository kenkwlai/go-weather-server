package cmd

import (
  "fmt"
  "log"

  "github.com/kenkwlai/weather-server/api"
  "github.com/kenkwlai/weather-server/app"
  "github.com/spf13/viper"
)

func InitApp() {
  a, err := app.New()
  if err != nil {
    log.Fatal(err)
  }
  defer a.Close()

  apis, err := api.New(a)
  if err != nil {
    log.Fatal(err)
  }

  InitServer(apis)
}

func InitConfig() {
  viper.SetConfigName("config")
  viper.AddConfigPath(".")
  viper.AddConfigPath("/etc/weather-server/")
  viper.AddConfigPath("$HOME/.weather-server")

  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err != nil {
    panic(fmt.Errorf("unable to read config: %s \n", err))
  }
}
