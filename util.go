package golisper

import "strings"

func tabulate(s string) string {
	arr := strings.Split(s, "\n")
	for i, ss := range arr {
		arr[i] = "    " + ss
	}
	return strings.Join(arr, "\n")
}
