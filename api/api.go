package api

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "os"
  "strings"
)

func Init(r *gin.Engine) {
  r.POST("/authorize", getToken)
  r.GET("/weather", jwtAuthenticator(), getWeather)
}

func jwtAuthenticator() gin.HandlerFunc {
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

      return []byte(os.Getenv("JWT_SECRET_KEY")), nil
    })

    switch err.(type) {
    case nil:
      if !token.Valid {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
      }

    case *jwt.ValidationError:
      vErr := err.(*jwt.ValidationError)
      switch vErr.Errors {
      case jwt.ValidationErrorExpired:
        c.String(http.StatusUnauthorized, "Token Expired")
        log.Println("Token Expired")
        return

      default:
        c.AbortWithStatus(http.StatusInternalServerError)
        log.Printf("ValidationError error: %+v\n", vErr.Errors)
        return
      }
    }

    if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
      if !claims.VerifyIssuer(os.Getenv("JWT_ISSUER"), true) {
        log.Println(err)
        c.AbortWithStatus(http.StatusUnauthorized)
        return
      }
    }
  }
}
