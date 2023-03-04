package main

import (
	"flag"
	"fmt"
	"time"
)

var overDaysFlag = flag.Int("overdays", 0, "calculate the average time spent over given days")

func main() {
	flag.Parse()
	inputs := flag.Args()
	overDays := *overDaysFlag

	if len(inputs) == 1 {
		input := inputs[0]
		duration, _ := time.ParseDuration(input)
		fmt.Print(duration)
		dailyAvg := dailyAverageMessage(overDays, duration)
		fmt.Println(dailyAvg)
		return
	}
	for _, input := range inputs {
		duration, _ := time.ParseDuration(input)
		fmt.Printf("%-10s = %9s\n", input, duration)
		dailyAvg := dailyAverageMessage(overDays, duration)
		fmt.Println(dailyAvg)
	}
}

func dailyAverageMessage(overDays int, duration time.Duration) (dailyAverage string) {
	if overDays > 0 {
		average := duration / time.Duration(overDays)
		average = average.Round(5 * time.Minute)
		dailyAverage = fmt.Sprintln("\tdaily average:", average)
	}
	return
}
