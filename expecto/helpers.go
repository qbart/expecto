package expecto

import (
	"io/fs"
	"os"
	"path/filepath"
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

func (s *Assert) NoErr(msg string, err error) {
	if err != nil {
		s.t.Logf("  %s:", msg)
		s.t.Logf("     got: \033[0;31m%v\033[0m", err)
		s.t.Fatalf("Error occurred")
	}
}

// TempFS expects args: "path", "content", "path2", "content2", ...
func TempFS(pathContentPair ...string) (fs.FS, string, func()) {
	dir, err := os.MkdirTemp("", "tempfs-*")
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		os.RemoveAll(dir)
	}

	for i := 1; i < len(pathContentPair); i += 2 {
		absPath := filepath.Join(dir, pathContentPair[i-1])
		filepathParts := strings.Split(pathContentPair[i-1], "/")
		if len(filepathParts) > 1 {
			dirsOnly := strings.Join(filepathParts[:len(filepathParts)-1], "/")
			if err := os.MkdirAll(filepath.Join(dir, dirsOnly), 0755); err != nil {
				panic(err)
			}
		}
		content := pathContentPair[i]
		if err := os.WriteFile(absPath, []byte(content), 0644); err != nil {
			panic(err)
		}
	}
	return os.DirFS(dir), dir, cleanup
}
