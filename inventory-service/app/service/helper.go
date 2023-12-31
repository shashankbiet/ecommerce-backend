package service

import "strings"

func setCategory(str *string) {
	*str = strings.ToUpper(strings.Join(strings.Split(*str, " "), "_"))
}

func setSubCategory(str *string) {
	*str = strings.ToUpper(strings.Join(strings.Split(*str, " "), "_"))
}
