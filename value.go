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
}

func NewVal(t ValueType, StringVal string, NumberVal float64, tag *Tag) *Value {
	return &Value{
		Type:      t,
		StringVal: StringVal,
		NumberVal: NumberVal,
		TagVal:    tag,
	}
}

func NewValNumber(n float64) *Value {
	return NewVal(TYPE_NUMBER, "", n, nil)
}

func NewValString(s string) *Value {
	return NewVal(TYPE_STRING, s, 0, nil)
}

func NewValEtc(s string) *Value {
	return NewVal(TYPE_ETC_STRING, s, 0, nil)
}

func NewValTag(t *Tag) *Value {
	return NewVal(TYPE_TAG, "", 0, t)
}

func (v *Value) String() string {
	if v.Type == TYPE_STRING {
		return "'" + v.StringVal + "'"
	} else if v.Type == TYPE_NUMBER {
		return fmt.Sprintf("%f", v.NumberVal)
	} else if v.Type == TYPE_ETC_STRING {
		return fmt.Sprintf("ETC(%s)", v.StringVal)
	} else if v.Type == TYPE_TAG {
		return v.TagVal.String()
	}
	return ""
}
