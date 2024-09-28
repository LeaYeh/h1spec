package config

import "testing"


// TestRunMode tests the RunMode method.
// e.g. http1.1/RFC7230/3/1 => means test protocal http1.1 in RFC7230 section 3, test case #1.
func TestRunMode(t *testing.T) {
	tests := []struct {
		sections []string
		target   string
		mode     int
	}{
		{sections: []string{"http1.1"}, target: "http1.1", mode: RunModeAll},
		{sections: []string{"http1.1"}, target: "http1.1/RFC7230", mode: RunModeAll},
		{sections: []string{"http1.1"}, target: "http1.1/RFC7230.2", mode: RunModeAll},
		{sections: []string{"http1.1"}, target: "http1.1/RFC7230.2.5", mode: RunModeAll},
		{sections: []string{"http1.1"}, target: "http1.1/RFC7230.2.5/1", mode: RunModeAll},

		{sections: []string{"http1.1/RFC7230"}, target: "http1.1", mode: RunModeGroup},
		{sections: []string{"http1.1/RFC7230"}, target: "http1.1/RFC7230", mode: RunModeAll},
		{sections: []string{"http1.1/RFC7230"}, target: "http1.1/RFC7230.2.5", mode: RunModeAll},
		{sections: []string{"http1.1/RFC7230"}, target: "http1.1/RFC7230.2.5/1", mode: RunModeAll},

		{sections: []string{"http1.1/RFC7230.2.5"}, target: "http1.1", mode: RunModeGroup},
		{sections: []string{"http1.1/RFC7230.2.5"}, target: "http1.1/RFC7230", mode: RunModeGroup},
		{sections: []string{"http1.1/RFC7230.2.5"}, target: "http1.1/RFC7230.2.5", mode: RunModeAll},
		{sections: []string{"http1.1/RFC7230.2.5"}, target: "http1.1/RFC7230.2.5/1", mode: RunModeAll},

		{sections: []string{"http1.1/RFC7230.2.5/1"}, target: "http1.1", mode: RunModeGroup},
		{sections: []string{"http1.1/RFC7230.2.5/1"}, target: "http1.1/RFC7230", mode: RunModeGroup},
		{sections: []string{"http1.1/RFC7230.2.5/1"}, target: "http1.1/RFC7230.2.5", mode: RunModeGroup},
		{sections: []string{"http1.1/RFC7230.2.5/1"}, target: "http1.1/RFC7230.2.5/1", mode: RunModeAll},
		{sections: []string{"http1.1/RFC7230.2.5/1"}, target: "http1.1/RFC7230.2.5/2", mode: RunModeNone},

		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.5", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.5.1", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.5.1.2", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.5.1.2/1", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.4", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.5.2", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.5.1.3", mode: RunModeAll},
		{sections: []string{"http1.1", "http1.1/RFC7230.5.1.2/1"}, target: "http1.1/RFC7230.5.1.2/2", mode: RunModeAll},
	}

	for i, tt := range tests {
		c := Config{
			Sections: tt.sections,
		}
		mode := c.RunMode(tt.target)
		if tt.mode != mode {
			t.Errorf("#%d mode - expect: %d, got: %d (%v / %v)", i, tt.mode, mode, tt.target, c.targetMap)
		}
	}
}
