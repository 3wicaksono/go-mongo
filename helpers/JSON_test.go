package helpers

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMakeJSON(t *testing.T) {
	Convey("Testing MakeJSON", t, func() {
		plainText := "macan ternak"

		json := MakeJSON(plainText)

		actual := false
		if json != "" {
			actual = true
		}
		Convey("Asserting return type", func() {
			So(actual, ShouldEqual, true)
		})
	})
}

func TestMakeJSONFail(t *testing.T) {
	Convey("Testing MakeJSON", t, func() {
		plainText := make(chan int)

		json := MakeJSON(plainText)

		actual := false
		if json == "" {
			actual = true
		}
		Convey("Asserting return type", func() {
			So(actual, ShouldEqual, true)
		})
	})
}
