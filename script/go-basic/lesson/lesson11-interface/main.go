package main

import (
	"fmt"

	"trungem.com/hoc-golang/cat"
	"trungem.com/hoc-golang/dog"
	"trungem.com/hoc-golang/mouse"
	"trungem.com/hoc-golang/service"
)

func MakeSound(a service.Animal) {
	fmt.Printf("Day la tieng cua %s: ", a.GetName())
	fmt.Println(a.Speak())
}

func MakeSoundPlus(p service.AnimalPlus) {
	fmt.Printf("Day la tieng cua %s: ", p.GetName())
	fmt.Println(p.Speak())
	fmt.Println(p.Eat())
}

func PrintValue(v interface{}) {
	//str, ok := v.(string)
	//if ok {
	//	fmt.Println(str)
	//} else {
	//	fmt.Println("Please send a string")
	//}
	//
	//numb, ok := v.(int)
	//if ok {
	//	fmt.Println(numb)
	//} else {
	//	fmt.Println("Please send a int")
	//}

	switch v.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("Type invalid")
	}
}

func main() {
	myDog, err := dog.New("   Bully   ")
	if err != nil {
		panic(err)
	}
	MakeSound(myDog)

	PrintValue("-=-=-=-=-=-=-=-=-=")

	myCat, err := cat.New("Pon")
	if err != nil {
		panic(err)
	}
	MakeSoundPlus(myCat)

	PrintValue("-=-=-=-=-=-=-=-=-=")

	myMouse, err := mouse.New("   Mr Ti   ")
	if err != nil {
		panic(err)
	}
	PrintValue(myMouse.Run())

	PrintValue(5)
	PrintValue(5.5)
	PrintValue(false)
}
