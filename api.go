package golisper

func Parse(src string) ([]*Value, error) {
	return parse(src)
}
