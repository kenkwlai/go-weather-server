package db

import (
  . "github.com/smartystreets/goconvey/convey"
  "github.com/spf13/viper"
  "testing"
)

func TestInitConfig(t *testing.T) {
  Convey("TestInitConfig", t, func() {
    Convey("TestInitConfigEmpty", func() {
      emptyConfig, err := InitConfig()
      So(err, ShouldNotBeNil)
      So(emptyConfig, ShouldBeNil)
    })

    Convey("TestInitConfigFailsWhenEitherOneIsEmpty", func() {
      viper.Set("DatabaseURI", "some value")
      viper.Set("DatabaseName", "")

      config, err := InitConfig()

      So(err, ShouldNotBeNil)
      So(config, ShouldBeNil)
    })

    Convey("TestInitConfigValid", func() {
      viper.Set("DatabaseURI", "some value")
      viper.Set("DatabaseName", "some value")

      config, err := InitConfig()

      So(err, ShouldBeNil)
      So(config.DatabaseURI, ShouldEqual, "some value")
      So(config.DatabaseName, ShouldEqual, "some value")
    })
  })
}
