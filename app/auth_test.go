package app

import (
  "github.com/kenkwlai/weather-server/models"
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

var credentials = models.UserCredentials{
  Username: "username",
  Password: "password",
}

func TestIssueJwt(t *testing.T) {
  Convey("TestIssueJwt", t, func() {
    Convey("TestIssueJwtSuccess", func() {
      result, err := IssueJwt(&credentials)

      So(err, ShouldBeNil)
      So(result, ShouldNotBeNil)
      So(result.Token, ShouldNotBeNil)
    })
  })
}
