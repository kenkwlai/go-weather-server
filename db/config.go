package db

import (
  "fmt"

  "github.com/spf13/viper"
)

type Config struct {
  DatabaseURI   string
  DatabaseName  string
}

func InitConfig() (*Config, error) {
  config := &Config{
    DatabaseURI: viper.GetString("DatabaseURI"),
    DatabaseName: viper.GetString("DatabaseName"),
  }
  if config.DatabaseURI == "" || config.DatabaseName == "" {
    return nil, fmt.Errorf("DatabaseURI and DatabaseName cannot be blank")
  }
  return config, nil
}
