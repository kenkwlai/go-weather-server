package app

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/kenkwlai/weather-server/config"
  "github.com/kenkwlai/weather-server/models"
  "time"
)

type Token struct {
  Expire   string   `json:"expire"`
  Token    string   `json:"token"`
}

const TimezoneLocationName = "Asia/Hong_Kong"
const TimeFormatLayout = "2006-01-02T15:04:05Z0700"

var jwtIssuer = config.GetOrDefault("JWT_ISSUER", "some-issuer")
var jwtSecretKey = config.GetOrDefault("JWT_SECRET_KEY", "some-secret-key")

func IssueJwt(user *models.UserCredentials) (*Token, error) {
  // sign the jwt token
  token := jwt.New(jwt.SigningMethodHS256)
  claims := make(jwt.MapClaims)
  loc, err := time.LoadLocation(TimezoneLocationName)
  if err != nil {
    return nil, err
  }

  now := time.Now().In(loc)
  exp := now.Add(time.Hour * time.Duration(1))
  claims["exp"] = exp.Unix()
  claims["iat"] = now.Unix()
  claims["iss"] = jwtIssuer
  claims["username"] = user.Username
  token.Claims = claims

  tokenString, err := token.SignedString([]byte(jwtSecretKey))
  if err != nil {
    return nil, err
  }

  return &Token{exp.Format(TimeFormatLayout), tokenString}, nil
}
