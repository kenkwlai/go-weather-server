package cmd

import (
  "github.com/kenkwlai/weather-server/db"
)

func InitApp() {
  defer db.Close()

  InitServer()
}
