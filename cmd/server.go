package cmd

import (
  "fmt"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/kenkwlai/weather-server/config"
  "log"
  "strconv"
  "time"

  "github.com/kenkwlai/weather-server/api"
)

func InitServer() {
  port, err := strconv.Atoi(config.GetOrDefault("PORT", "8000"))

  if err != nil {
    log.Fatal(err)
  }

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

  router.Run(fmt.Sprintf(":%v", port))
}
