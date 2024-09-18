package config

import "testing"


// TestRunMode tests the RunMode method.
// e.g. http1.1/RFC7230/3/1 => means test protocal http1.1 in RFC7230 section 3, test case #1.
func TestRunMode(t *testing.T) {
	tests := []struct {
		section string
		want    int
	}{
		{"", RunModeNone},
		{"a", RunModeAll},
		{"a/b", RunModeGroup},
		{"a/b/c", RunModeNone},
		{"a/1", RunModeAll},
		{"a/1/2", RunModeGroup},
	}
}
