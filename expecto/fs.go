package expecto

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

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
