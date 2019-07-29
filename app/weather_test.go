package app

import (
  "fmt"
  "github.com/kenkwlai/weather-server/models"
  "reflect"
  "testing"

  . "github.com/bouk/monkey"
  "github.com/kenkwlai/weather-server/db"
  . "github.com/smartystreets/goconvey/convey"
)

var validWeatherData = models.CurrentWeatherData{
  Coord: models.Coord{
    Lat: 1.0,
    Lon: 1.0,
  },
  Weather: []models.Weather{
    {
      Id: 1,
      Main: "some string",
      Description: "some string",
      Icon: "some string",
    },
  },
  // Not setting other fields here...
}

func TestGetWeatherByCity(t *testing.T) {
  Convey("TestGetWeatherByCity", t, func() {
    Convey("GetWeather", func() {
      var mockDb *db.Database
      guard := PatchInstanceMethod(reflect.TypeOf(mockDb), "GetWeather", func(_ *db.Database, _ string) (*models.CurrentWeatherData, error) {
        return &validWeatherData, nil
      })
      defer guard.Unpatch()

      output, err := mockDb.GetWeather("any")
      So(output, ShouldEqual, &validWeatherData)
      So(err, ShouldBeNil)
    })

    Convey("GetWeatherFailed", func() {
      var mockDb *db.Database
      guard := PatchInstanceMethod(reflect.TypeOf(mockDb), "GetWeather", func(_ *db.Database, _ string) (*models.CurrentWeatherData, error) {
        return nil, fmt.Errorf("some error")
      })
      defer guard.Unpatch()

      output, err := mockDb.GetWeather("any")
      So(output, ShouldBeNil)
      So(err, ShouldNotBeNil)
    })
  })
}

func TestCreateWeather(t *testing.T) {
  Convey("TestDatabaseCreateWeather", t, func() {
    Convey("CreateWeather", func() {
      var mockDb *db.Database
      guard := PatchInstanceMethod(reflect.TypeOf(mockDb), "CreateWeather", func(_ *db.Database, _ *models.CurrentWeatherData) error {
        return nil
      })
      defer guard.Unpatch()

      err := mockDb.CreateWeather(&validWeatherData)
      So(err, ShouldBeNil)
    })

    Convey("CreateWeatherFailed", func() {
      var mockDb *db.Database
      guard := PatchInstanceMethod(reflect.TypeOf(mockDb), "CreateWeather", func(_ *db.Database, _ *models.CurrentWeatherData) error {
        return fmt.Errorf("some error")
      })
      defer guard.Unpatch()

      err := mockDb.CreateWeather(&validWeatherData)
      So(err, ShouldNotBeNil)
    })
  })
}

