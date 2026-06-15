package array

import "fmt"

type Nhanvien struct {
	Id   int
	Name string
	Age  int
}

func Array() {
	employees := [...]Nhanvien{
		{Id: 1, Name: "Alice", Age: 30},
		{Id: 2, Name: "Bob", Age: 20},
		{Id: 3, Name: "Charlie", Age: 20},
		{Id: 4, Name: "David", Age: 20},
	}

	//fmt.Println(employees[1].Id)
	//fmt.Println(employees[1].Name)
	//fmt.Println(employees[1].Age)

	for _, val := range employees {
		fmt.Printf("Name: %s, Age: %d\n", val.Name, val.Age)
	}

	//var number int
	//fmt.Println(number)
	//
	//var numbers [5]int
	//fmt.Println(numbers)
	//
	//var character string
	//fmt.Println(character)
	//
	//var characters [5]string
	//fmt.Println(characters)
	//
	//var numbers [5]int
	//numbers[2] = 10 // [0 0 10 0 0]
	//fmt.Println(numbers)
	//
	//var numbers = [5]int{1, 3, 5}
	//numbers[3] = 10
	//fmt.Println(numbers)
	//
	//var numbers = [...]int{4, 76, 9}
	//fmt.Printf("Total array: %T\n", numbers)
	//fmt.Println(numbers)

	//var matrix = [2][3]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//}
	//matrix[1][1] = 7
	//
	//fmt.Println(matrix)
	//fmt.Println(matrix[1][1])

	//numbers := [5]int{1, 2, 3, 4, 5}
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}

	//numbers := [3][4]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//}
	//
	//for i := 0; i < len(numbers); i++ {
	//	for j := 0; j < len(numbers[i]); j++ {
	//		fmt.Println(numbers[i][j])
	//	}
	//}

	//numbers := [5]int{1, 2, 3, 4, 5}
	//
	//for i, val := range numbers {
	//	fmt.Println(i, val)
	//}

	//numbers := [3][4]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//}
	//
	//for _, number := range numbers {
	//	for _, element := range number {
	//		fmt.Println(element)
	//	}
	//}
}
