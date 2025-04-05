package expecto

import (
	"reflect"
	"strings"
	"testing"
)

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	return &Assert{t: t}
}

func (s *Assert) Eq(msg string, value, expected any) {
	if !reflect.DeepEqual(value, expected) {
		s.t.Logf("  %s:", msg)
		s.t.Logf("    want: \033[0;32m%v\033[0m", expected)
		s.t.Logf("     got: \033[0;31m%v\033[0m", value)
		s.t.Fatalf("Values do not match")
	}
}

func (s *Assert) Nil(msg string, value any) {
	reflectValue := reflect.ValueOf(value)
	if !reflectValue.IsNil() {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", value)
		s.t.Fatalf("Value is not nil")
	}
}

func (s *Assert) NotNil(msg string, value any) {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.IsNil() {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", value)
		s.t.Fatalf("Value is nil")
	}
}

func (s *Assert) True(msg string, value bool) {
	if !value {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", value)
		s.t.Fatalf("Value is not true")
	}
}

func (s *Assert) False(msg string, value bool) {
	if value {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", value)
		s.t.Fatalf("Value is not false")
	}
}

func (s *Assert) Contains(msg string, value, expected string) {
	if !strings.Contains(value, expected) {
		s.t.Logf("  %s:", msg)
		s.t.Logf("    want: \033[0;32m%v\033[0m", expected)
		s.t.Logf("     got: \033[0;31m%v\033[0m", value)
		s.t.Fatalf("Value does not contain expected")
	}
}
