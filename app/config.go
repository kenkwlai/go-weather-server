package app

import (
  "fmt"

  "github.com/spf13/viper"
)

type Config struct {
  SecretKey []byte
  Issuer    string
}

func InitConfig() (*Config, error) {
  config := &Config{
    SecretKey: []byte(viper.GetString("SecretKey")),
    Issuer:    viper.GetString("Issuer"),
  }

  if len(config.SecretKey) == 0 {
    return nil, fmt.Errorf("SecretKey must be set")
  }

  if config.Issuer == "" {
    return nil, fmt.Errorf("issuer must be set")
  }

  return config, nil
}
