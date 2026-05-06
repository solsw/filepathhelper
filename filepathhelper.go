package filepathhelper

import (
	"math/rand"
	"path/filepath"
	"strings"

	"github.com/solsw/pathhelper"
)

// Random returns a random name (for directory or file).
func Random() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	const randLen = 10
	b := make([]byte, randLen)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// WithoutExt returns 'path' without extension and corresponding dot.
// If 'path' ends with dot, the dot is removed.
// If 'path' has no extension or is empty, 'path' is returned.
func WithoutExt(path string) string {
	if path == "" {
		return ""
	}
	e := filepath.Ext(path)
	if e == "" {
		return path
	}
	return path[:len(path)-len(e)]
}

// ChangeExt changes extension of 'path' to 'ext'. 'ext' may or may not start with dot.
// If 'path' or 'ext' is empty or 'ext' is ".", 'path' is returned.
func ChangeExt(path, ext string) string {
	if path == "" || ext == "" || ext == "." {
		return path
	}
	if strings.HasPrefix(ext, ".") {
		return WithoutExt(path) + ext
	}
	return WithoutExt(path) + "." + ext
}

// Split splits 'path' (using [filepath.Separator] as separator)
// into slice of strings containing directory names and filename.
// (E.g. on Windows "a\b\c.d" is split into {"a", "b", "c.d"} slice.)
func Split(path string) []string {
	return pathhelper.Split(filepath.ToSlash(path))
}

// StartSeparator returns 'path' guaranteed to start with [filepath.Separator].
// If 'path' is empty, empty string is returned.
func StartSeparator(path string) string {
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, string(filepath.Separator)) {
		return path
	}
	return string(filepath.Separator) + path
}

// EndSeparator returns 'path' guaranteed to end with [filepath.Separator].
// If 'path' is empty, empty string is returned.
func EndSeparator(path string) string {
	if path == "" {
		return ""
	}
	if strings.HasSuffix(path, string(filepath.Separator)) {
		return path
	}
	return path + string(filepath.Separator)
}
