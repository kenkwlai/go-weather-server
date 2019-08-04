package cmd

import (
  "github.com/kenkwlai/weather-server/db"
)

func InitApp() {
  db.Init()
  defer db.Close()

  InitServer()
}
