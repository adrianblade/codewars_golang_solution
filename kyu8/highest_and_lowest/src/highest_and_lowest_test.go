package src

import "testing"

func TestHighAndLow(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"", "8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighAndLow(tt.args); got != tt.want {
				t.Errorf("HighAndLow() = %v, want %v", got, tt.want)
			}
		})
	}
}