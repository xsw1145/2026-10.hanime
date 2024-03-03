package util

import "strings"

var InvalidDirSymbols = [10]rune{'/', '<', '>', ':', '"', '/', '\\', '|', '?', '*'}

func ReplaceChars(str string, chars []rune) string {
	var strs []string
	for _, c := range chars {
		strs = append(strs, string(c), "")
	}
	r := strings.NewReplacer(strs...)
	return r.Replace(str)
}
