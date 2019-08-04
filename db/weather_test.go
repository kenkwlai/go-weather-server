package db

import (
  . "bou.ke/monkey"
  "fmt"
  . "github.com/golang/mock/gomock"
  . "github.com/kenkwlai/weather-server/db/mocks"
  . "github.com/kenkwlai/weather-server/models"
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

func TestWeatherMongoStore(t *testing.T) {
  Convey("TestWeatherMongoStore", t, func() {
    Convey("TestWeatherMongoStoreShouldReturnAnInstance", func() {
      var mockInstance *store

      patch := Patch(WeatherMongoStore, func() *store {
        return mockInstance
      })
      defer patch.Unpatch()

      result := WeatherMongoStore()

      So(result, ShouldEqual, mockInstance)
    })

    Convey("TestWeatherMongoStoreShouldReturnNil", func() {
      patch := Patch(WeatherMongoStore, func() *store {
        return nil
      })
      defer patch.Unpatch()

      result := WeatherMongoStore()

      So(result, ShouldBeNil)
    })
  })
}

func TestStore_GetWeather(t *testing.T) {
  Convey("TestGetWeather", t, func() {
    Convey("TestGetWeatherFromStoreSuccess", func() {
      ctrl := NewController(t)
      defer ctrl.Finish()

      m := NewMockWeatherStore(ctrl)

      m.
        EXPECT().
        GetWeather(Any()).
        Return(&weatherResult, nil)

      result, err := m.GetWeather("HK")

      So(err, ShouldBeNil)
      So(result, ShouldNotBeNil)
      So(result, ShouldEqual, &weatherResult)
    })

    Convey("TestGetWeatherFromStoreFailsShouldReturnError", func() {
      var errMsg = "some error"

      ctrl := NewController(t)
      defer ctrl.Finish()

      m := NewMockWeatherStore(ctrl)

      m.
        EXPECT().
        GetWeather(Any()).
        Return(nil, fmt.Errorf(errMsg))

      result, err := m.GetWeather("HK")

      So(result, ShouldBeNil)
      So(err, ShouldNotBeNil)
      So(err.Error(), ShouldEqual, errMsg)
    })
  })
}

func TestStore_CreateWeather(t *testing.T) {
  Convey("TestCreateWeather", t, func() {
    Convey("TestCreateWeatherInStoreSuccess", func() {
      ctrl := NewController(t)
      defer ctrl.Finish()

      m := NewMockWeatherStore(ctrl)

      m.
        EXPECT().
        CreateWeather(Any()).
        Return(nil)

      err := m.CreateWeather(&weatherResult)

      So(err, ShouldBeNil)
    })

    Convey("TestCreateWeatherInStoreFailsShouldReturnError", func() {
      var errMsg = "some error"

      ctrl := NewController(t)
      defer ctrl.Finish()

      m := NewMockWeatherStore(ctrl)

      m.
        EXPECT().
        CreateWeather(Any()).
        Return(fmt.Errorf(errMsg))

      err := m.CreateWeather(&weatherResult)

      So(err, ShouldNotBeNil)
      So(err.Error(), ShouldEqual, errMsg)
    })
  })
}
