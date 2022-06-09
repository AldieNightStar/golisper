package main

import (
	"fmt"

	"github.com/AldieNightStar/golisper"
)

func main() {
	tags, _ := golisper.Parse(
		"(1 2 3) (4 5 6) (9 10 (a b))",
	)
	fmt.Println(tags)
}
