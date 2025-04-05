package expecto

import (
	"strings"
)

func (s *Assert) NoErr(msg string, err error) {
	if err != nil {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", err)
		s.t.Fatalf("Error occurred")
	}
}

func (s *Assert) Err(msg string, err error) {
	if err == nil {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", err)
		s.t.Fatalf("Expected error")
	}
}

func (s *Assert) ErrContains(msg string, err error, substr string) {
	if err == nil || !strings.Contains(err.Error(), substr) {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", err)
		s.t.Fatalf("Expected error containing %s", substr)
	}
}
