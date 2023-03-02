package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	inputs := os.Args[1:]
	if len(inputs) == 1 {
		input := inputs[0]
		duration, _ := time.ParseDuration(input)
		fmt.Println(duration)
		return
	}
	for _, input := range inputs {
		duration, _ := time.ParseDuration(input)
		fmt.Printf("%-10s = %9s\n", input, duration)
	}
}
