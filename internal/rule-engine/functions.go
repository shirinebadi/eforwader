package ruleengine

import (
	"regexp"
	"strings"
)

func Is(a, b string) bool {
	return a == b
}

func IsNot(a, b string) bool {
	return a != b
}

func Contains(a, b string) bool {
	return strings.Contains(a, b)
}

func DoesNotContain(a, b string) bool {
	return !strings.Contains(a, b)
}

func BeginsWith(a, b string) bool {
	return strings.HasPrefix(a, b)
}

func EndsWith(a, b string) bool {
	return strings.HasSuffix(a, b)
}

func IsEmpty(a string) bool {
	return len(a) == 0
}

func IsNotEmpty(a string) bool {
	return len(a) > 0
}

func IsGroupMemberIn(a string, group []string) bool {
	for _, member := range group {
		if a == member {
			return true
		}
	}
	return false
}

func IsNotGroupMemberIn(a string, group []string) bool {
	return !IsGroupMemberIn(a, group)
}

func Matches(a, pattern string) bool {
	matched, _ := regexp.MatchString(pattern, a)
	return matched
}

func DoesNotMatch(a, pattern string) bool {
	return !Matches(a, pattern)
}
