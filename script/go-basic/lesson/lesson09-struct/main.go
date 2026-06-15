package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GiangVien struct {
	Name   string `json:"hoten"`
	Gender int    `json:"gioitinh"`
	Phone  string `json:"dienthoai"`
}

func (gv *GiangVien) hienThiThongTin() {
	fmt.Printf("Ho ten: %s \n", gv.Name)
	fmt.Printf("Gioi tinh: %d \n", gv.Gender)
	fmt.Printf("Dien thoai: %s \n", gv.Phone)
}

func (gv *GiangVien) clear() {
	gv.Name = ""
	gv.Gender = 0
	gv.Phone = ""
}

func main() {

	trungem := GiangVien{
		Name:   "trungem",
		Gender: 1,
		Phone:  "091233213",
	}

	//trungem.hienThiThongTin()
	//
	//trungem.clear()
	//
	//fmt.Println()
	//
	//trungem.hienThiThongTin()

	output, err := json.Marshal(trungem)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}
