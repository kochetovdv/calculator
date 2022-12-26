package errorchecking

import (
	"regexp"
)

// RegexpPatternCheck checks is string matched with pattern. True if match was found
func RegexpPatternCheck(pattern, s string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
