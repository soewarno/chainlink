package web

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func MatchWildcardBoxPath(boxList []string, path string, file string) (matchedPath string) {
	idPattern := regexp.MustCompile(`(_[a-zA-Z0-9]+_)`)
	pathSeparator := (string)(os.PathSeparator)
	escapedPathSeparator := pathSeparator
	if pathSeparator == `\` {
		escapedPathSeparator = `\\`
	}
	pathAndFile := filepath.Clean(strings.Join(
		[]string{path, file},
		pathSeparator,
	))
	normalizedPathAndFile := strings.Replace(
		strings.TrimPrefix(pathAndFile, `/`),
		`/`,
		pathSeparator,
		-1,
	)

	for i := 0; i < len(boxList) && matchedPath == ""; i++ {
		boxPathWithIDPattern := idPattern.ReplaceAllString(boxList[i], `[a-zA-Z0-9]+`)
		pathPattern := fmt.Sprintf(
			`^%s$`,
			strings.Replace(boxPathWithIDPattern, `\`, escapedPathSeparator, -1),
		)
		match, _ := regexp.MatchString(pathPattern, normalizedPathAndFile)

		if match {
			matchedPath = boxList[i]
		}
	}

	return matchedPath
}

func MatchExactBoxPath(boxList []string, path string) (matchedPath string) {
	pathWithoutPrefix := strings.TrimPrefix(path, "/")

	for i := 0; i < len(boxList) && matchedPath == ""; i++ {
		boxPath := boxList[i]
		if boxPath == pathWithoutPrefix {
			matchedPath = boxPath
		}
	}

	return matchedPath
}
