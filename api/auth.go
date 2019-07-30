package api

import (
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/kenkwlai/weather-server/models"
  "net/http"
)

func (api *API) GetToken(c *gin.Context) {
  var credentials models.UserCredentials

  err := json.NewDecoder(c.Request.Body).Decode(&credentials)
  if err != nil {
    c.String(http.StatusUnauthorized, "Error when getting token: %v", err)
    return
  }

  // Some trivial validation
  if credentials.Username == "" || credentials.Password == "" {
    c.String(http.StatusUnauthorized, "Unauthorized credentials")
    return
  }

  token, err := api.App.IssueJwt(&credentials)
  if err != nil {
    c.Error(err)
    c.String(http.StatusInternalServerError, "internal server error")
    return
  }

  c.JSON(http.StatusOK, token)
}
