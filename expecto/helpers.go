package expecto

import (
	"reflect"
	"strings"
	"testing"
)

func Eq(t *testing.T, msg string, value, expected any) {
	if !reflect.DeepEqual(value, expected) {
		t.Logf("  %s:", msg)
		t.Logf("    want: \033[0;32m%v\033[0m", expected)
		t.Logf("     got: \033[0;31m%v\033[0m", value)
		t.Fatalf("Values do not match")
	}
}

func Nil(t *testing.T, msg string, value any) {
	reflectValue := reflect.ValueOf(value)
	if !reflectValue.IsNil() {
		t.Logf("  %s:", msg)
		t.Logf("     got: \033[0;31m%v\033[0m", value)
		t.Fatalf("Value is not nil")
	}
}

func NotNil(t *testing.T, msg string, value any) {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.IsNil() {
		t.Logf("  %s:", msg)
		t.Logf("     got: \033[0;31m%v\033[0m", value)
		t.Fatalf("Value is nil")
	}
}

func True(t *testing.T, msg string, value bool) {
	if !value {
		t.Logf("  %s:", msg)
		t.Logf("     got: \033[0;31m%v\033[0m", value)
		t.Fatalf("Value is not true")
	}
}

func False(t *testing.T, msg string, value bool) {
	if value {
		t.Logf("  %s:", msg)
		t.Logf("     got: \033[0;31m%v\033[0m", value)
		t.Fatalf("Value is not false")
	}
}

func Contains(t *testing.T, msg string, value, expected string) {
	if !strings.Contains(value, expected) {
		t.Logf("  %s:", msg)
		t.Logf("    want: \033[0;32m%v\033[0m", expected)
		t.Logf("     got: \033[0;31m%v\033[0m", value)
		t.Fatalf("Value does not contain expected")
	}
}
