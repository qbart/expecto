package expecto

import (
	"strings"
	"testing"
)

func NoErr(t *testing.T, msg string, err error) {
	if err != nil {
		t.Logf("  %s:", msg)
		t.Logf("     got: \033[0;31m%v\033[0m", err)
		t.Fatalf("Error occurred")
	}
}

func Err(t *testing.T, msg string, err error) {
	if err == nil {
		t.Logf("  %s:", msg)
		t.Logf("     got: \033[0;31m%v\033[0m", err)
		t.Fatalf("Expected error")
	}
}

func ErrContains(t *testing.T, msg string, err error, substr string) {
	Err(t, msg, err)

	if !strings.Contains(err.Error(), substr) {
		t.Logf("  %s:", msg)
		t.Logf("    want: \033[0;32m%v\033[0m", substr)
		t.Logf("     got: \033[0;31m%v\033[0m", err)
		t.Fatalf("Expected error containing %s", substr)
	}
}
