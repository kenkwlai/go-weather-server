package config

import (
  "log"
  "os"
)

func GetOrDefault(envKey string, defaultValue string) string {
  env := os.Getenv(envKey)
  if env == "" {
    log.Printf("Env variable not found for key: %v", envKey)
    return defaultValue
  }

  return env
}
