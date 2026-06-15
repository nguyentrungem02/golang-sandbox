package main

import "fmt"

func main() {
	s1 := 15
	s2 := 4

	phepsosanh := s1 <= s2
	fmt.Printf("Ket qua so sanh: %t \n", phepsosanh)

	phepluanly := true && true && true
	fmt.Printf("Ket qua luan ly: %t \n", !phepluanly)

	//tong := s1 + s2
	//hieu := s1 - s2
	//tich := s1 * s2
	//thuong := float64(s1) / float64(s2)
	//chialaydu := s1 % s2
	//s1 /= 4
	//
	//fmt.Printf("Tong cua %d va %d la %d \n", s1, s2, tong)
	//fmt.Printf("Hieu cua %d va %d la %d \n", s1, s2, hieu)
	//fmt.Printf("Tich cua %d va %d la %d \n", s1, s2, tich)
	//fmt.Printf("Thuong cua %d va %d la %.2f \n", s1, s2, thuong)
	//fmt.Printf("Chia lay du cua %d va %d la %d \n", s1, s2, chialaydu)
	//fmt.Printf("S1 = %d \n", s1)
}
