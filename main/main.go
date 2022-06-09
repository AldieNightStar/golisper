package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AldieNightStar/golisper"
)

func main() {
	file, _ := os.Open("file.txt")
	defer file.Close()
	fileBytes, _ := ioutil.ReadAll(file)
	fileContent := string(fileBytes)
	tags, err := golisper.Parse(fileContent)

	for _, tag := range tags {
		fmt.Println(tag)
	}
	fmt.Println(err)
}
