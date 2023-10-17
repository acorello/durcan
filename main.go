package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const (
	divByK      = "div-by"
	truncateToK = "truncate-to"
)

var divByFlag = flag.Int(divByK, 0, "divide each value by given number")
var truncateToFlag = flag.String(truncateToK, "1s", "truncate to given amount (eg. `5m` means `5 minutes`), default 1s")

func main() {
	flag.Parse()
	truncateTo := *truncateToFlag
	divBy := *divByFlag

	truncateToDuration, err := time.ParseDuration(truncateTo)
	if err != nil {
		log.Fatalf("Invalid -%s value %q; error: %v", truncateToK, truncateTo, err)
	}

	Run(flag.Args(), divBy, truncateToDuration)
}

func Run(inputs []string, divBy int, truncateToDuration time.Duration) {
	divByMsg := fmt.Sprintf("%s %2d", divByK, divBy)
	for _, input := range inputs {
		duration, err := time.ParseDuration(input)
		if err != nil {
			fmt.Printf("error: invalid input %q: %v", input, err)
			continue
		}
		rounded := duration.Round(truncateToDuration)
		fmt.Printf("%-11s = %9s\n", input, rounded)
		if divBy > 0 {
			average := duration / time.Duration(divBy)
			average = average.Truncate(truncateToDuration)
			fmt.Printf("  %-9s = %9s\n", divByMsg, average)
		}
	}
}
