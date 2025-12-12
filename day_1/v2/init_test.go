package day1

import "testing"

func Test_turn(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		p                byte
		rangeC           int
		expectedPassword int
	}{
		// TODO: Add test cases.
		{
			p:                'L',
			rangeC:           50,
			name:             "Check for exect reset to zero having final password to be 1",
			expectedPassword: 1,
		},
		{
			p:                'L',
			rangeC:           100,
			name:             "Check for one complete round coming back to 50 after passing by zero having final password to be 1",
			expectedPassword: 2,
		},
	}
	for _, tt := range tests {
		initialPos := 50
		finalPassword := 0

		t.Run(tt.name, func(t *testing.T) {
			turn(tt.p, tt.rangeC, &initialPos, &finalPassword)
		})
	}
}
