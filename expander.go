package main

import (
	"regexp"
	"strings"
	"time"
)

// Exapnd file name by replacing date/time info using time.Now().Format()
// IDEA: Add support for other expanders, such as custom variables, locations, system info, etc.
func expandFileName(filename string) string {
	re := regexp.MustCompile("{([a-zA-Z0-9.:-_]+)}")

	t := time.Now()

	return re.ReplaceAllStringFunc(filename, func(part string) string {
		return t.Format(
			strings.TrimSuffix(
				strings.TrimPrefix(part, "{"),
				"}",
			),
		)
	})
}
