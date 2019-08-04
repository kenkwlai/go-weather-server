package config

import (
  . "github.com/smartystreets/goconvey/convey"
  "os"
  "testing"
)

func TestGetOrDefault(t *testing.T) {
  Convey("TestGetOrDefault", t, func() {
    Convey("TestGetOrDefaultWithDefault", func() {
      envKey := "some-key"
      configVal := "abc"

      result := GetOrDefault(envKey, configVal)

      So(result, ShouldEqual, configVal)
    })

    Convey("TestGetOrDefaultWithEnv", func() {
      envKey := "some-key"
      envVal := "some-value"
      configVal := "abc"

      os.Setenv(envKey, envVal)

      result := GetOrDefault(envKey, configVal)

      So(result, ShouldEqual, envVal)
    })
  })
}
