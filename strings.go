package emoji

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func StringToAliases(s string) string {
	s = EmojisToAliases(s)

	needsFixing := false

	for _, r := range s {
		if utf8.RuneLen(r) > 3 {
			needsFixing = true
		}
	}

	if !needsFixing {
		return s
	}

	var b strings.Builder

	for _, r := range s {
		if utf8.RuneLen(r) > 3 {
			b.WriteString(fmt.Sprintf(":u%4x:", int32(r)))
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}
