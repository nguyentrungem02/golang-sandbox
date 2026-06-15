package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Bai 1
	//for i := 1; i <= 100; i++ {
	//	if i == 6 || i == 48 || i == 75 || i == 89 {
	//		continue
	//	}
	//	if i == 100 {
	//		fmt.Printf("%d", i)
	//		continue
	//	}
	//	fmt.Printf("%d, ", i)
	//}
	//fmt.Println()

	// Bai 2
	//var count = 1
	//for j := 1; j <= 100; j++ {
	//	if j%2 == 0 {
	//		continue
	//	}
	//
	//	count++
	//
	//	if count == 4 {
	//		count = 1
	//		fmt.Printf("%d", j)
	//		fmt.Println()
	//		continue
	//	}
	//	if j == 99 {
	//		fmt.Printf("%d", j)
	//		continue
	//	}
	//	fmt.Printf("%d, ", j)
	//}
	//
	//fmt.Println()

	// Bai 3
	var start, end int
	fmt.Print("Enter your start: ")
	fmt.Scanf("%d", &start)

	fmt.Print("Enter your end: ")
	fmt.Scanf("%d", &end)

	if start == 0 || end == 0 {
		fmt.Println("Number start and end different 0")
	} else if start > end {
		fmt.Println("Number end is greater than start")
	} else {
		xhtml := ""
		for i := start; i <= end; i++ {
			xhtml += "Bang cuu chuong " + strconv.Itoa(i) + "\n"

			for j := 1; j <= 10; j++ {
				xhtml += strconv.Itoa(i) + " x " + strconv.Itoa(j) + " = " + strconv.Itoa(i*j) + "\n"
			}
		}
		fmt.Println(xhtml)
	}
}
