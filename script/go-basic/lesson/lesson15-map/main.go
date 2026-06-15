package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	//employee := map[string]Employee{
	//	"e1": {Name: "e1", Age: 18, Role: "Admin"},
	//	"e2": {Name: "e2", Age: 19, Role: "User"},
	//}
	//
	//fmt.Printf("%+v \n", employee)
	//
	//for _, v := range employee {
	//	fmt.Printf("Name: %s \n", v.Name)
	//	fmt.Printf("Age: %d \n", v.Age)
	//	fmt.Printf("Role: %s \n", v.Role)
	//
	//}

	studentSubject := map[string][]string{
		"Trung Em": {"Toan", "Golang"},
		"Dang":     {"CNTT", "Vat ly"},
	}

	fmt.Println(studentSubject)

	fmt.Println(studentSubject["Trung Em"])
	fmt.Println(studentSubject["Trung Em"][0])

	for k, v := range studentSubject {
		for _, vv := range v {

			fmt.Printf("Mon hoc cua %s la %s\n", k, vv)
		}
	}

	//// Cach 1
	//drink := map[string]int{
	//	"tea":    500,
	//	"coffee": 1000,
	//}
	//
	//fmt.Println(drink)
	//
	//student := map[int]string{
	//	10: "Tuan",
	//	20: "Trung Em",
	//}
	//
	//fmt.Println(student)
	//
	//// Cach 2
	//m := make(map[string]int)
	//m["tea"] = 500
	//m["coffee"] = 1000
	//fmt.Println(m)
	//
	//// Cach 3
	//var monan map[string]int
	//monan = make(map[string]int)
	//monan["chao"] = 500
	//monan["com"] = 1000
	//monan["bun"] = 900
	//monan["ga"] = 300
	//fmt.Println(monan)
	//
	//// Kiem tra co ton tai khong
	//value, exists := monan["chao"]
	//if exists {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("Khong ton tai trong map")
	//}
	//
	//// Duyet map
	//for key, value := range monan {
	//	fmt.Println(key, value)
	//}
}
