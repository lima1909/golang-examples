package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetFilePrefix_Dir(t *testing.T) {
	dirPrefix := getFilePrefix(true)
	t.Logf("FirPrefix: %s", dirPrefix)

	Convey(`Compare dirPrefix with "d "`, t, func() {
		So(dirPrefix, ShouldEqual, "d ")
	})
}

func TestGetFilePrefix_File(t *testing.T) {
	filePrefix := getFilePrefix(false)
	t.Logf("FilePrefix: %s", filePrefix)
	Convey(`Compare filePrefix with "- "`, t, func() {
		So(filePrefix, ShouldEqual, "- ")
	})
}
