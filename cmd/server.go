package cmd

import (
  "fmt"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "time"

  "github.com/kenkwlai/weather-server/api"
)

func InitServer(api *api.API) {

  corsOption := cors.New(cors.Config{
    AllowAllOrigins:  true,
    AllowMethods:     []string{"GET", "HEAD", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
  })

  router := gin.Default()
  router.Use(corsOption)
  api.Init(router)

  router.Run(fmt.Sprintf(":%d", api.Config.Port))
}
