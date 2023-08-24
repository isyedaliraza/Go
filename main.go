package main

import (
	"fmt"
	"learn-go/src/iteration"
)

func main() {
	output := iteration.Repeat("a", 3)
	fmt.Println(output)
}
