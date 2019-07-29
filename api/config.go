package api

import "github.com/spf13/viper"

type Config struct {
  Port              int
  OpenWeatherApiKey string
}

func InitConfig() (*Config, error) {
  config := &Config{
    Port: viper.GetInt("Port"),
    OpenWeatherApiKey: viper.GetString("OpenWeatherApiKey"),
  }
  if config.Port == 0 {
    config.Port = 8080
  }
  return config, nil
}
