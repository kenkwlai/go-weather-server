package api

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  "github.com/kenkwlai/weather-server/app"
  "log"
  "net/http"
  "strings"
)

type API struct {
  App    *app.App
  Config *Config
}

func New(a *app.App) (api *API, err error) {
  api = &API{App: a}
  api.Config, err = InitConfig()
  if err != nil {
    return nil, err
  }

  return api, nil
}

func (api *API) Init(r *gin.Engine) {
  r.POST("/authorize", api.GetToken)
  r.GET("/weather", api.JwtAuthenticator(), api.GetWeather)
}

func (api *API) JwtAuthenticator() gin.HandlerFunc {
  return func(c *gin.Context) {
    tokenStr := c.GetHeader("Authorization")
    if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer") {
      c.AbortWithStatus(http.StatusUnauthorized)
      return
    }

    jwtToken := strings.Split(tokenStr, " ")[1]
    token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (i interface{}, e error) {
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
      }

      return []byte(api.App.Config.SecretKey), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
      if !claims.VerifyIssuer(api.App.Config.Issuer, true) {
        log.Println(err)
        c.AbortWithStatus(http.StatusUnauthorized)
        return
      }
    }
  }
}
