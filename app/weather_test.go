package app

import (
  "fmt"
  . "github.com/kenkwlai/weather-server/models"
  . "github.com/prashantv/gostub"
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

var weatherResult = CurrentWeatherData{
  Coord:      Coord{},
  Weather:    nil,
  Base:       "",
  Main:       Main{},
  Visibility: 0,
  Wind:       Wind{},
  Clouds:     Cloud{},
  Rain:       Rain{},
  Snow:       Snow{},
  Dt:         0,
  Sys:        Sys{},
  Timezone:   0,
  Id:         0,
  Name:       "HK",
  Cod:        0,
}

func TestGetWeatherByCity(t *testing.T) {
  Convey("TestGetWeatherByCity", t, func() {
    Convey("TestGetWeatherByCitySuccess", func() {
      fun := GetWeatherByCity

      stubs := Stub(&fun, func(_ string) (*CurrentWeatherData, error) {
        return &weatherResult, nil
      })
      defer stubs.Reset()

      result, err := fun("HK")

      So(err, ShouldBeNil)
      So(result, ShouldNotBeNil)
      So(result, ShouldEqual, &weatherResult)
    })

    Convey("TestGetWeatherByCityFailsShouldReturnError", func() {
      var errMsg = "some error"

      fun := GetWeatherByCity

      stubs := Stub(&fun, func(_ string) (*CurrentWeatherData, error) {
        return nil, fmt.Errorf(errMsg)
      })
      defer stubs.Reset()

      result, err := fun("HK")

      So(result, ShouldBeNil)
      So(err, ShouldNotBeNil)
      So(err.Error(), ShouldEqual, errMsg)
    })
  })
}

func TestCreateWeather(t *testing.T) {
  Convey("TestCreateWeather", t, func() {
    Convey("TestCreateWeatherSuccess", func() {
      fun := CreateWeather

      stubs := Stub(&fun, func(_ *CurrentWeatherData) error {
        return nil
      })
      defer stubs.Reset()

      err := fun(&weatherResult)

      So(err, ShouldBeNil)
    })

    Convey("TestCreateWeatherFailsShouldReturnError", func() {
      var errMsg = "some error"

      fun := CreateWeather

      stubs := Stub(&fun, func(_ *CurrentWeatherData) error {
        return fmt.Errorf(errMsg)
      })
      defer stubs.Reset()

      err := fun(&weatherResult)

      So(err, ShouldNotBeNil)
      So(err.Error(), ShouldEqual, errMsg)
    })
  })
}
