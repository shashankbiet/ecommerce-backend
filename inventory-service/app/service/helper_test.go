package service

import "testing"

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "trims and uppercases", input: " mobile phone ", want: "MOBILE_PHONE"},
		{name: "preserves repeated spaces as underscores", input: "mobile  phone", want: "MOBILE__PHONE"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.input
			normalizeName(&got)
			if got != test.want {
				t.Errorf("normalizeName() = %q, want %q", got, test.want)
			}
		})
	}
}
