package golisper

import (
	"fmt"
	"strings"
)

type Tag struct {
	Name   string
	Values []*Value
	Line   int
}

func NewTag(name string, vals []*Value, line int) *Tag {
	return &Tag{Name: name, Values: vals, Line: line}
}

func (t *Tag) String() string {
	sb := strings.Builder{}
	sb.Grow(32)
	sb.WriteString("TAG[")
	sb.WriteString(t.Name)
	sb.WriteString("] LINE:")
	sb.WriteString(fmt.Sprintf("%d", t.Line))
	sb.WriteString(" (")
	for _, v := range t.Values {
		sb.WriteString("\n")
		sb.WriteString(tabulate(v.String()))
	}
	sb.WriteString("\n)")
	return sb.String()
}
