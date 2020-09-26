package util

import (
	"bytes"
	"unicode/utf8"
)

// RunIndex ...
func RunIndex(s string, edge int) int {
	step := 0
	idx := 0
	for i, w := 0, 0; i < len(s); i += w {
		step++
		_, w = utf8.DecodeRuneInString(s[i:])
		if step == edge {
			idx = i + w
			break
		}
	}
	return idx
}

// SplitSubN ...
func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
