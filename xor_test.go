package xor

import "testing"

func TestXOR(t *testing.T) {
	tests := []struct {
		inputs   []bool
		expected bool
	}{
		{
			[]bool{true, false, false, false},
			true,
		},
		{
			[]bool{false, false, false, true},
			true,
		},
		{
			[]bool{true, false, true, false},
			false,
		},
		{
			[]bool{true, false, true, true},
			false,
		},
		{
			[]bool{true, true, true, true},
			false,
		},
		{
			[]bool{false, false, false, false},
			false,
		},
		{
			[]bool{true, true, true, false},
			false,
		},
		{
			[]bool{true},
			true,
		},
		{
			[]bool{false},
			false,
		},
	}

	for _, test := range tests {
		actual := xor(test.inputs...)
		if actual != test.expected {
			t.Errorf("unexpected value: got %t, but expected %t\n", actual, test.expected)
		}
	}
}
