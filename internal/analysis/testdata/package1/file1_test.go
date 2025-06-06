package package1_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		expected string
	}{
		"sub test name": {
			expected: "expected",
		},
		"sub test name includes (regexp meta characters)": {
			expected: "expected",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := "actual"
			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Errorf("\n(-expected, +actual)\n%s", diff)
			}
		})
	}
}

func TestBar(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		expected string
	}{
		"sub test name": {
			expected: "expected",
		},
		"sub test name includes (regexp meta characters)": {
			expected: "expected",
		},
	} {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := "actual"
			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Errorf("\n(-expected, +actual)\n%s", diff)
			}
		})
	}
}

func TestBaz(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "slice based sub test name",
			expected: "expected",
		},
		{
			name:     "slice based sub test name with (regexp meta characters)",
			expected: "expected",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := "actual"
			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Errorf("\n(-expected, +actual)\n%s", diff)
			}
		})
	}
}
