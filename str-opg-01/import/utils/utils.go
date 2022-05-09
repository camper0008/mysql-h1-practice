package utils

import (
	"strings"
)

func DanishToEnglishFieldNames(names []string) []string {
	res := []string{}
	for i := 0; i < len(names); i++ {
		newStr := names[i]
		newStr = strings.Replace(newStr, "æ", "ae", -1)
		newStr = strings.Replace(newStr, "ø", "oe", -1)
		newStr = strings.Replace(newStr, "å", "aa", -1)
		res = append(res, newStr)
	}
	return res
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
