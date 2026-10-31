package service

import "strings"

func normalizeName(str *string) {
	*str = strings.ToUpper(strings.Join(strings.Split(strings.TrimSpace(*str), " "), "_"))
}
