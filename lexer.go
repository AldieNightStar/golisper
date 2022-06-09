package golisper

import "fmt"

func lex(src string) ([]*token, *GoLisperError) {
	pos := 0
	line := 1
	toks := make([]*token, 0, 32)
	for pos < len(src) {
		chr := src[pos]
		if chr == ' ' || chr == '\t' || chr == '\r' {
			pos += 1
			continue
		}
		if chr == '\n' {
			line += 1
			pos += 1
			continue
		}
		com, comCnt := lexComment(src[pos:])
		if comCnt > 0 {
			toks = append(toks, newToken(tokenComment, com, 0, line))
			pos += comCnt
			continue
		}
		sym, symCnt := lexSymbol(src[pos:])
		if symCnt > 0 {
			toks = append(toks, newToken(tokenSymbol, sym, 0, line))
			pos += symCnt
			continue
		}
		str, strCnt := lexString(src[pos:])
		if strCnt > 0 {
			toks = append(toks, newToken(tokenString, str, 0, line))
			pos += strCnt
			continue
		}
		num, numCnt := lexNumber(src[pos:])
		if numCnt > 0 {
			toks = append(toks, newToken(tokenNumber, "", num, line))
			pos += numCnt
			continue
		}
		etc, etcCnt := lexEtc(src[pos:])
		if etcCnt > 0 {
			toks = append(toks, newToken(tokenEtc, etc, 0, line))
			pos += etcCnt
			continue
		}
		return nil, newError(ErrUnknownToken, fmt.Sprintf("Unkown token '%s ...'. Line: %d", src[pos:pos+10], line))
	}
	return toks, nil
}

type tokenType int

const (
	tokenString tokenType = iota
	tokenNumber
	tokenSymbol
	tokenEtc
	tokenComment
)

type token struct {
	typ  tokenType
	val  string
	num  float64
	line int
}

func newToken(t tokenType, val string, num float64, line int) *token {
	return &token{typ: t, val: val, num: num, line: line}
}
