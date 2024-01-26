package util

import (
	"regexp"
	"strings"
	"unicode"
)

const sqlReg = "\\b(and|exec|insert|select|drop|grant|alter|delete|update|count|chr|mid|master|truncate|char|declare|or)\\b|(\\*|;|\\+|'|%|--)"
const normalStrReg = "^[a-zA-Z0-9_-]+$"
const intReg = "^[0-9]+$"

func IsNormalStr(str string) bool {
	isNormalStr, err := regexp.MatchString(normalStrReg, str)
	if err != nil {
		panic(err)
	}
	isSql, err := regexp.MatchString(sqlReg, strings.ToLower(str))
	if err != nil {
		panic(err)
	}
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
	isInt, err := regexp.MatchString(intReg, str)
	if err != nil {
		panic(err)
	}
	return isInt
}
