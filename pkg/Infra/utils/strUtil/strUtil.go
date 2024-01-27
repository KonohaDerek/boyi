package strUtil

import "strings"

func ContainArray(str string, s []string) bool {
	for i := range s {
		if strings.Contains(str, s[i]) {
			return true
		}
	}
	return false
}
