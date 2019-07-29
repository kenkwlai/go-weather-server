package app

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

    Convey("TestInitConfigFailsWhenEmpty", func() {
      viper.Set("SecretKey", "")

      config, err := InitConfig()

      So(err, ShouldNotBeNil)
      So(config, ShouldBeNil)
    })

    Convey("TestInitConfigValid", func() {
      viper.Set("SecretKey", "some value")

      config, err := InitConfig()

      So(err, ShouldBeNil)
      So(config.SecretKey, ShouldNotBeNil)
    })
  })
}
