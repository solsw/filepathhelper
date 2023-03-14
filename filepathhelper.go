package filepathhelper

import (
	"path/filepath"

	"github.com/solsw/pathhelper"
)

// NoExt returns 'path' without extension.
// If 'path' has no extension or is empty, 'path' is returned.
func NoExt(path string) string {
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
// If 'path' or 'ext' is empty, 'path' is returned.
func ChangeExt(path, ext string) string {
	if path == "" || ext == "" || ext == "." {
		return path
	}
	if ext[0] == '.' {
		return NoExt(path) + ext
	}
	return NoExt(path) + "." + ext
}

// Split splits 'path' (using [filepath.Separator] as separator)
// into slice of strings containing directory names and filename.
// (E.g. on Windows "a\b\c.d" is splitted into {"a", "b", "c.d"} slice.)
func Split(path string) []string {
	return pathhelper.Split(filepath.ToSlash(path))
}

// StartSeparator returns 'path' guaranteed to start with [filepath.Separator].
// If 'path' is empty, empty string is returned.
func StartSeparator(path string) string {
	if path == "" {
		return ""
	}
	if path[0] == filepath.Separator {
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
	if path[len(path)-1] == filepath.Separator {
		return path
	}
	return path + string(filepath.Separator)
}
