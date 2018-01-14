package cmd

import "testing"

func TestGetStartDirectoryFromArgs(t *testing.T) {
	var dirTestData = []struct {
		args        []string
		expectedDir string
	}{
		{nil, "./"},
		{[]string{"mydir"}, "mydir"},
		{[]string{"mydir", "yourdir"}, "mydir"},
	}

	for _, row := range dirTestData {
		dir := getStartDirectoryFromArgs(row.args)
		t.Logf("Find dir: %v for args: %v", dir, row.args)
		if dir != row.expectedDir {
			t.Errorf("Expected: %v, but was %v", row.expectedDir, dir)
		}

	}
}
