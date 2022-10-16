package commandparser

import "strings"

func ExtractRnameFromTag(s string) string {
	s = strings.Trim(s, "`")
	strs := strings.Split(s, " ")
	for _, str := range strs {
		if strings.HasPrefix(str, "rname:") {
			return strings.Trim(strings.TrimPrefix(str, "rname:"), "\"")
		}
	}
	return ""
}
