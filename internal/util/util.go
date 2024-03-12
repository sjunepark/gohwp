package util

import "strings"

type Stringer interface {
	String() string
}

func JoinStringers[T Stringer](slice []T, delimiter string) string {
	var builder strings.Builder
	for i, s := range slice {
		if i > 0 {
			builder.WriteString(delimiter)
		}
		builder.WriteString(s.String())
	}
	return builder.String()
}
