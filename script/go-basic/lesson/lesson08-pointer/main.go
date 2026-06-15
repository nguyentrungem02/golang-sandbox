package main

import "fmt"

func lythuyetPointer() {
	name := "Trung Em"

	fmt.Println("-=-=-=-=-= Information name variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %s \n", name)
	fmt.Printf("Address: %v \n", &name)

	// Tao pointer
	fmt.Println()
	ptrName := &name
	fmt.Println("-=-=-=-=-= Information ptrName variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", ptrName)
	fmt.Printf("Value: %v \n", ptrName)
	fmt.Printf("Address: %v \n", &ptrName)

	fmt.Printf("Find value name from ptrName %v \n", *ptrName)

	// Tao pointer
	fmt.Println()
	ptrName2 := &ptrName
	fmt.Println("-=-=-=-=-= Information ptrName2 variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", ptrName2)
	fmt.Printf("Value: %v \n", ptrName2)
	fmt.Printf("Address: %v \n", &ptrName2)

	fmt.Printf("Find value ptrName from ptrName2 %v \n", *ptrName2)
	fmt.Printf("Find value name from ptrName2 %v \n", **ptrName2)
}

func updateName(name string) {
	name = "TEm nè"
	fmt.Printf("Value: %s \n", name)
}

func main() {
	name := "Trung Em"
	fmt.Println("-=-=-=-=-= Information name variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %s \n", name)
	fmt.Printf("Address: %v \n", &name)

	fmt.Println()
	updateName(name)

	fmt.Println("-=-=-=-=-= Information name variable after run update value -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %s \n", name)
	fmt.Printf("Address: %v \n", &name)
}
