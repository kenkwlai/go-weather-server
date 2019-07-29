package app

import (
  "context"
  "github.com/kenkwlai/weather-server/db"
)

type App struct {
  Config *Config
  Database *db.Database
}

func New() (app *App, err error) {
  app = &App{}
  app.Config, err = InitConfig()
  if err != nil {
    return nil, err
  }

  dbConfig, err := db.InitConfig()
  if err != nil {
    return nil, err
  }

  app.Database, err = db.Connect(dbConfig)
  if err != nil {
    return nil, err
  }

  return app, err
}

func (a *App) Close() error {
  return a.Database.Client().Disconnect(context.Background())
}
