package golisper

func Parse(src string) ([]*Tag, error) {
	return parse(src)
}
