package utils

import (
	"fmt"
	"testing"
)

// AssertEqual compares expected and actual boolean values and reports a test failure if they are not equal.
func AssertEqual(t *testing.T, expected, actual bool) {
	if expected != actual {
		t.Fatalf(`
Not equal!
expected: %v
actual:   %v`,
			expected,
			actual,
		)
	}
}

// TestName formats an integer and a string into a combined string with an underscore separator.
func TestName(i int, name string) string {
	return fmt.Sprintf("%d_%s", i, name)
}
