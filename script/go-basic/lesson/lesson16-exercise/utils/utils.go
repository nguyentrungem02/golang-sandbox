package utils

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

		fmt.Println("❌ Giá trị không hợp lệ! Hãy nhập số nguyên dương.")
	}
}

func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}

		fmt.Println("❌ Không được bỏ trống! Hãy nhập ít nhất một ký tự.")
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
		fmt.Println("❌ Lỗi xoá màn hình")
	}
}

func GenerateId() string {
	id := uuid.New()
	return id.String()
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
