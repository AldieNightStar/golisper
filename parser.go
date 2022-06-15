package golisper

import "fmt"

func parse(src string) ([]*Value, error) {
	toks, err := lex(src)
	if err != nil {
		return nil, err
	}
	pos := 0
	values := make([]*Value, 0, 32)
	for pos < len(toks) {
		tok := toks[pos]
		if tok.typ == tokenComment {
			// Ignore the comments
			pos += 1
			continue
		}
		if tok.typ != tokenSymbol {
			return nil, newError(ErrUnknownValue, fmt.Sprintf("Unkown value on %d line. Should be tag here", tok.line))
		}
		tag, tagCnt, err := parseTag(toks[pos:])
		if err != nil {
			return nil, err
		}
		if tagCnt > 0 {
			values = append(values, NewValTag(tag, tag.Line))
			pos += tagCnt
			continue
		}
		return nil, newError(ErrUnknownValue, fmt.Sprintf("It's not a tag. Line: %d", tok.line))
	}
	return values, nil
}

func parseValue(tok *token, line int) (val *Value) {
	if tok.typ == tokenNumber {
		return NewValNumber(tok.num, line)
	} else if tok.typ == tokenString {
		return NewValString(tok.val, line)
	} else if tok.typ == tokenEtc {
		return NewValEtc(tok.val, line)
	}
	return nil
}

func parseTag(toks []*token) (tag *Tag, count int, err error) {
	first := toks[0]
	if first.typ != tokenSymbol || first.val != "(" {
		return nil, 0, nil
	}
	values := make([]*Value, 0, 8)
	pos := 1
	for pos < len(toks) {
		tok := toks[pos]
		if tok.typ == tokenSymbol && tok.val == ")" {
			pos += 1
			break
		}
		val := parseValue(tok, tok.line)
		if val != nil {
			values = append(values, val)
			pos += 1
			continue
		}
		tag, tagCnt, err := parseTag(toks[pos:])
		if err != nil {
			return nil, 0, err
		}
		if tagCnt > 0 {
			values = append(values, NewValTag(tag, tok.line))
			pos += tagCnt
			continue
		}
		if tok.typ == tokenComment {
			// Comments are ignored
			pos += 1
			continue
		}
		return nil, 0, newError(ErrTagValueErr, fmt.Sprintf("Wrong tag at line: %d", tok.line))
	}
	if len(values) < 1 {
		return nil, 0, newError(ErrTagValueErr, fmt.Sprintf("Line %d: Empty tag error. No values", first.line))
	}
	if values[0].Type != TYPE_ETC_STRING {
		return nil, 0, newError(ErrUnknownToken, fmt.Sprintf("Line %d: Tag name is wrong. No values", values[0].Line))
	}
	tagName := values[0].StringVal
	return NewTag(tagName, values[1:], values[0].Line), pos, nil
}
