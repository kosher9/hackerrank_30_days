package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var ii uint64
	var dd float64
	var ss string

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	ii,_ = strconv.ParseUint(scanner.Text(), 0, 64)
	scanner.Scan()
	dd,_ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	ss = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i+ ii)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n",d+dd)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s+ ss)
}
