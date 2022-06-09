package golisper

import (
	"strconv"
	"strings"
)

const digits = "0123456789"
const symbols = "()[]<>{}:;\\/!@#$%^&*_+-=~?.,|"

func lexSymbol(s string) (sym string, count int) {
	if strings.IndexByte(symbols, s[0]) != -1 {
		return string(s[0]), 1
	}
	return "", 0
}

func lexString(s string) (tok string, count int) {
	if len(s) < 2 {
		return "", 0
	}
	if !(s[0] == '\'' || s[0] == '"' || s[0] == '`') {
		return "", 0
	}
	ptr := 1
	end := rune(s[0])
	esc := false
	sb := strings.Builder{}
	sb.Grow(64)
	for _, c := range s[1:] {
		if esc {
			esc = false
			if c == 'n' {
				sb.WriteRune('\n')
			} else if c == 't' {
				sb.WriteRune('\t')
			} else if c == 'r' {
				sb.WriteRune('\r')
			} else if c == '0' {
				sb.WriteRune(0)
			} else {
				sb.WriteRune(c)
			}
			ptr += 1
			continue
		}
		if c == '\\' {
			ptr += 1
			esc = true
			continue
		}
		if c == '\n' {
			return "", 0
		}
		if c == end {
			ptr += 1
			break
		}
		ptr += 1
		sb.WriteRune(c)
	}
	return sb.String(), ptr
}

func lexNumber(s string) (result float64, count int) {
	sb := strings.Builder{}
	sb.Grow(8)
	dotAllow := true
	ptr := 0
	for i, c := range s {
		if c == '.' && i == 0 && dotAllow {
			sb.WriteString("0.")
			dotAllow = false
			ptr += 1
		} else if c == '-' && i == 0 {
			ptr += 1
			sb.WriteRune('-')
		} else if strings.Contains(digits, string(c)) {
			sb.WriteRune(c)
			ptr += 1
		} else if c == '.' && dotAllow {
			dotAllow = false
			ptr += 1
			sb.WriteRune('.')
		} else {
			break
		}
	}
	if n, err := strconv.ParseFloat(sb.String(), 64); err == nil {
		return n, ptr
	} else {
		return 0, 0
	}
}

func lexEtc(s string) (tok string, count int) {
	sb := strings.Builder{}
	sb.Grow(32)
	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' || strings.ContainsRune("()[]<>{}", c) {
			break
		}
		sb.WriteRune(c)
	}
	return sb.String(), sb.Len()
}
