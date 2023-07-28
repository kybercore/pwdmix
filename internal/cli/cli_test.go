package cli

import "testing"

func TestParseFlags(t *testing.T) {

	// Test with default flags
	_, err := ParseFlags()
	if err != nil {
		t.Error("Unexpected error for default flags")
	}
}
