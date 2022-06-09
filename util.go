package golisper

import (
	"io/ioutil"
	"os"
	"strings"
)

func tabulate(s string) string {
	arr := strings.Split(s, "\n")
	for i, ss := range arr {
		arr[i] = "    " + ss
	}
	return strings.Join(arr, "\n")
}

func LoadFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(fileBytes), nil
}
