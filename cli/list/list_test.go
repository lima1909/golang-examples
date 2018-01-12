package list

import (
	"testing"
)

func TestGetFilePrefix(t *testing.T) {
	dirPrefix := getFilePrefix(true)
	t.Logf("FirPrefix: %s", dirPrefix)
	if dirPrefix != "d " {
		t.Error("Expected \"d \", but was ", dirPrefix)
	}

	filePrefix := getFilePrefix(false)
	t.Logf("FilePrefix: %s", filePrefix)
	if filePrefix != "- " {
		t.Error("Expected \"- \", but was ", filePrefix)
	}
}
