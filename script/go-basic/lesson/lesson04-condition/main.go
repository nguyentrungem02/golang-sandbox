package main

import "fmt"

func main() {
	//diem := 7
	//
	//if diem >= 8 {
	//	fmt.Println("Diem lon hon 8")
	//} else {
	//	fmt.Println("Diem nho hon 8")
	//}
	//
	//toan := 12
	//if toan > 8 {
	//	fmt.Println("Hoc sinh gioi")
	//} else if toan <= 8 && toan >= 6 {
	//	fmt.Println("Hoc sinh kha")
	//} else if toan <= 5 && toan >= 3 {
	//	fmt.Println("Hoc sinh trung binh")
	//} else {
	//	fmt.Println("Hoc sinh yeu")
	//}
	//
	//fmt.Println("Out")

	//monan := "bun"
	//
	//switch monan {
	//case "com", "bun":
	//	fmt.Println("Mon nay la com")
	//case "chao", "my":
	//	fmt.Println("Mon nay la chao")
	//case "pho", "lau":
	//	fmt.Println("Mon nay la pho")
	//default:
	//	fmt.Println("Khong an gi ca")
	//
	//}
	//i := 1
	//for i < 10 {
	//	fmt.Printf("Number %d \n", i)
	//	i++
	//}
	//
	//for j := 1; j <= 100; j++ {
	//	fmt.Printf("Number %d \n", j)
	//
	//}

	// break
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			//break
			continue
		}
		fmt.Printf("%d \n", i)
	}
}
