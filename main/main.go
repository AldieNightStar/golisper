package main

import (
	"fmt"

	"github.com/AldieNightStar/golisper"
)

func main() {
	fileString, _ := golisper.LoadFile("file.lsp")
	tags, err := golisper.Parse(fileString)

	for _, tag := range tags {
		fmt.Println(tag)
	}
	fmt.Println(err)
}
