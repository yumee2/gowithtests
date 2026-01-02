package iteration

import "strings"

func Repeat(str string, times int) string {
	var repeated strings.Builder

	for range times {
		repeated.WriteString(str)
	}

	return repeated.String()
}
