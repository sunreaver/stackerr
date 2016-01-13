package stackerr

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_New(t *testing.T) {
	Convey("Test StackError New()", t, func() {

		Convey("param are string", func() {
			err := New("Something Wrong")
			So(err.Filename, ShouldEqual, "stackerr_test.go")
			So(err.Line, ShouldEqual, 14)
			So(err.ErrorMessage, ShouldEqual, "Something Wrong")
			So(err.Detail(), ShouldEqual, "{stackerr_test.go:14} Something Wrong")
			//Println(err.Detail())
		})

		Convey("param are StackErr", func() {
			err1 := New("Something Wrong")
			err2 := New(err1)
			So(err1, ShouldEqual, err2)
		})

		Convey("param are StackErr", func() {
			err := errors.New("Something Wrong")
			stackErr := New(err)
			So(stackErr.Filename, ShouldEqual, "stackerr_test.go")
			So(stackErr.Line, ShouldEqual, 30)
			So(stackErr.ErrorMessage, ShouldEqual, "Something Wrong")
		})
	})
}
