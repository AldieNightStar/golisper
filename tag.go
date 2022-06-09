package golisper

import "strings"

type Tag struct {
	Name   string
	Values []*Value
}

func NewTag(name string, vals []*Value) *Tag {
	return &Tag{Name: name, Values: vals}
}

func (t *Tag) String() string {
	sb := strings.Builder{}
	sb.Grow(32)
	sb.WriteString("TAG(")
	sb.WriteString(t.Name)
	sb.WriteString(" ")
	for _, v := range t.Values {
		sb.WriteString(v.String())
	}
	sb.WriteString(")")
	return sb.String()
}
