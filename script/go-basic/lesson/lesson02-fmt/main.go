package main

import "fmt"

func main() {
	//fmt.Print("Hello Trung Em")
	//fmt.Print("Hello Trung Em")

	//fmt.Println("Hello Trung Em")
	//fmt.Println("Hello Trung Em")

	//var fullName = "Trung Em"
	//var age = 20
	//fmt.Printf("Hello %s and your age is %d .\n", fullName, age)

	//var firstName, lastName string
	//fmt.Print("Enter full name: ")
	//fmt.Scan(&firstName, &lastName)
	//fmt.Println("Hello " + firstName + " " + lastName)

	//var firstName, lastName string
	//fmt.Print("Enter first name: ")
	//fmt.Scanln(&firstName)
	//fmt.Print("Enter last name: ")
	//fmt.Scanln(&lastName)
	//fmt.Println("Hello " + firstName + " " + lastName)

	//var name string
	//var age int
	//fmt.Print("Enter your name: ")
	//fmt.Scanf("%s", &name)
	//fmt.Print("Enter your age: ")
	//fmt.Scanf("%d", &age)
	//fmt.Printf("My name is %s and age is %d \n ", name, age)

	//message := fmt.Sprint("My name is ", "Trung Em")
	//fmt.Println(message)

	//message := fmt.Sprintln("My name is ", "Trung Em")
	//fmt.Println(message)

	//name := "Trung Em"
	//age := 20
	//message := fmt.Sprintf("My name is %s and age is %d", name, age)
	//fmt.Println(message)

	ten := "Trung Em"
	tuoi := 20
	chieucao := 1.56
	daTotNghiep := true
	phanTram := 10

	fmt.Printf("Kiểu dữ liệu của biến ten: %T \n", ten)                 // Kiểu dữ liệu của biến ten: string
	fmt.Printf("Kiểu dữ liệu của biến tuoi: %T \n", tuoi)               // Kiểu dữ liệu của biến tuoi: int
	fmt.Printf("Kiểu dữ liệu của biến chieucao: %T \n", chieucao)       // Kiểu dữ liệu của biến chieucao: float64
	fmt.Printf("Kiểu dữ liệu của biến daTotNghiep: %T \n", daTotNghiep) // Kiểu dữ liệu của biến daTotNghiep: bool
	fmt.Printf("Kiểu dữ liệu của biến phanTram: %T \n", phanTram)       // Kiểu dữ liệu của biến phanTram: int

	fmt.Printf("Tôi tên là: %v \n", ten) // Tôi tên là: Trung Em

	fmt.Printf("Tôi tên là %s và tôi %d tuổi \n", ten, tuoi) // Tôi tên là Trung Em và tôi 20 tuổi

	fmt.Printf("Chieu cao của tôi là %.2f \n", chieucao) // Chieu cao của tôi là 1.56
	fmt.Printf("Chieu cao của tôi là %.5f \n", chieucao) // Chieu cao của tôi là 1.56000
	fmt.Printf("Chieu cao của tôi là %.1f \n", chieucao) // Chieu cao của tôi là 1.6

	fmt.Printf("Tôi đã tốt nghiệp: %t \n", daTotNghiep) // Tôi đã tốt nghiệp: true

	fmt.Printf("Tôi đã học %d%% \n", phanTram) // Tôi đã học 10%
}
