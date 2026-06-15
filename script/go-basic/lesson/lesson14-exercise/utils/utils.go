package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		val, err := strconv.Atoi(input)
		if err == nil && val > 0 {
			return val
		}

		fmt.Println("❌ Gia tri khong hop le! Hay nhap so nguyen duong.")
	}
}

func GetPositiveFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		val, err := strconv.ParseFloat(input, 64)
		if err == nil && val > 0 {
			return val
		}

		fmt.Println("❌ Gia tri khong hop le! Hay nhap so thuc nguyen duong.")
	}
}

func GetOptionalPositiveFloat(prompt string, oldValue float64) float64 {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}

	val, err := strconv.ParseFloat(input, 64)
	if err != nil && val < 0 {
		fmt.Println("❌ Gia tri khong hop le! Giu nguyen gia tri cu.")
		return oldValue
	}

	return val
}

func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}

		fmt.Println("❌ Gia tri khong hop le! Hay nhap it nhat mot ky tu.")
	}
}

func GetOptionalString(prompt string, oldValue string) string {
	for {
		input := ReadInput(prompt)
		if input == "" {
			return oldValue
		}

		return input
	}
}

func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error clear screen")
	}
}

type HasId interface {
	GetId() int
}

func IsIdUnique[T HasId](id int, list []T) bool {
	for _, student := range list {
		if student.GetId() == id {
			return false
		}
	}
	return true
}
