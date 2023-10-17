package main

import "time"

var args = []string{"3600s", "6s", "1s"}

func ExampleRun_noDivBy_rounding5sec() {
	Run(args, 0, 5*time.Second)
	//OUTPUT:
	// 3600s       =    1h0m0s
	// 6s          =        5s
	// 1s          =        0s
}

func ExampleRun_divBy2_rounding5sec() {
	Run(args, 2, 5*time.Second)
	//OUTPUT:
	// 3600s       =    1h0m0s
	//   div-by  2 =     30m0s
	// 6s          =        5s
	//   div-by  2 =        0s
	// 1s          =        0s
	//   div-by  2 =        0s
}
