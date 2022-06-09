package main

import (
	"fmt"

	"github.com/AldieNightStar/golisper"
)

func main() {
	tags, err := golisper.Parse(
		"(nums 1 2 3) (nums 4 5 6) (nums 9 10 (a b))",
	)
	for _, tag := range tags {
		fmt.Println(tag)
	}
	fmt.Println(err)
}
