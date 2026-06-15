package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Box[T any] struct {
	Content     T
	Description T
}

func PrintValue[Type any](value Type) {
	fmt.Println(value)
}

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

func IsNotEqual[T comparable](a, b T) bool {
	return a != b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](a, b T) T {
	return a + b
}

func MaxLengthString(a, b string) string {
	if len(a) > len(b) {
		return a
	}

	return b
}

func main() {
	//PrintValue("Trung Em")
	//PrintValue(10)
	//PrintValue(true)

	//result01 := IsEqual(0, 1)
	//PrintValue(result01)
	//
	//result02 := IsEqual("Tem", "Em")
	//PrintValue(result02)
	//
	//result03 := IsEqual(6.7, 2.4)
	//PrintValue(result03)
	//
	//result04 := IsNotEqual(6.7, 2.4)
	//PrintValue(result04)

	//fmt.Println(Max(1, 2))
	//fmt.Println(Max(5.5, 9))
	//fmt.Println(Max("Em", "Tem"))
	//fmt.Println(MaxLengthString("Em", "Tem"))

	//stringBox := Box[string]{
	//	Content:     "Hoc golang Generic",
	//	Description: "Mo ta hoc golang",
	//}
	//
	//intBox := Box[int]{
	//	Content:     123,
	//	Description: 123,
	//}
	//
	//PrintValue(stringBox.Content)
	//PrintValue(stringBox.Description)
	//PrintValue(intBox.Content)
	//PrintValue(intBox.Description)

	PrintValue(Sum(5, 3.4))
	PrintValue(Sum(5.5, 3.4))
	PrintValue(Sum(5, 3))
}
