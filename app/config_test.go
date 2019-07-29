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

    Convey("TestInitConfigFailsWhenSecretKeyIsEmpty", func() {
      viper.Set("SecretKey", "")
      viper.Set("Issuer", "some value")

      config, err := InitConfig()

      So(err, ShouldNotBeNil)
      So(config, ShouldBeNil)
    })

    Convey("TestInitConfigFailsWhenIssuerIsEmpty", func() {
      viper.Set("SecretKey", "some value")
      viper.Set("Issuer", "")

      config, err := InitConfig()

      So(err, ShouldNotBeNil)
      So(config, ShouldBeNil)
    })

    Convey("TestInitConfigValid", func() {
      viper.Set("SecretKey", "some value")
      viper.Set("Issuer", "some value")

      config, err := InitConfig()

      So(err, ShouldBeNil)
      So(config.SecretKey, ShouldNotBeNil)
    })
  })
}
