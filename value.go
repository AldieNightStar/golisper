package golisper

import "fmt"

type ValueType int

const (
	TYPE_NUMBER ValueType = iota
	TYPE_STRING
	TYPE_ETC_STRING
	TYPE_TAG
)

type Value struct {
	Type      ValueType
	StringVal string
	NumberVal float64
	TagVal    *Tag
	Line      int
}

func NewVal(t ValueType, StringVal string, NumberVal float64, tag *Tag, line int) *Value {
	return &Value{
		Type:      t,
		StringVal: StringVal,
		NumberVal: NumberVal,
		TagVal:    tag,
		Line:      line,
	}
}

func NewValNumber(n float64, line int) *Value {
	return NewVal(TYPE_NUMBER, "", n, nil, line)
}

func NewValString(s string, line int) *Value {
	return NewVal(TYPE_STRING, s, 0, nil, line)
}

func NewValEtc(s string, line int) *Value {
	return NewVal(TYPE_ETC_STRING, s, 0, nil, line)
}

func NewValTag(t *Tag, line int) *Value {
	return NewVal(TYPE_TAG, "", 0, t, line)
}

func (v *Value) String() string {
	if v.Type == TYPE_STRING {
		return "STR:'" + v.StringVal + "' "
	} else if v.Type == TYPE_NUMBER {
		return fmt.Sprintf("%f ", v.NumberVal)
	} else if v.Type == TYPE_ETC_STRING {
		return fmt.Sprintf("ETC(%s) ", v.StringVal)
	} else if v.Type == TYPE_TAG {
		return v.TagVal.String()
	}
	return ""
}
