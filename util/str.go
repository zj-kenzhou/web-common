package util

import (
	"regexp"
	"strings"
	"unicode"
)

var sqlReg = regexp.MustCompile("\\b(and|exec|insert|select|drop|grant|alter|delete|update|count|chr|mid|master|truncate|char|declare|or)\\b|(\\*|;|\\+|'|%|--)")
var normalStrReg = regexp.MustCompile("^[a-zA-Z0-9_-]+$")
var intReg = regexp.MustCompile("^[0-9]+$")

func IsNormalStr(str string) bool {
	isNormalStr := normalStrReg.MatchString(str)
	isSql := sqlReg.MatchString(strings.ToLower(str))
	return isNormalStr && !isSql
}

// CameCaseToUnderscore 驼峰转换下划线
func CameCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

func IsInt(str string) bool {
	isInt := intReg.MatchString(str)
	return isInt
}
