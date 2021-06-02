package day2

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBasic(t *testing.T) {
	username := "boy"
	fmt.Printf("welcome, %s", username)
}

func TestNumber(t *testing.T) {
	a := 255
	fmt.Printf("二进制: %b\n", a)
	fmt.Printf("八进制: %o\n", a)
	fmt.Printf("十进制: %d\n", a)
	fmt.Printf("十六进制: %x\n", a)
	fmt.Printf("大写十六进制: %X\n", a)

	fmt.Printf("十六进制: %d\n", Hex2Dec("4E2D"))
	fmt.Printf("字符: %c\n", 20013)
	fmt.Printf("Unicode格式: %U\n", '中') // U+4E2D
}

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

func Hex2Bin(val string) string {
	return fmt.Sprintf("%b", Hex2Dec(val))
}

func Bin2Hex(val string) string {
	ui, err := strconv.ParseUint(val, 2, 64)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%x", ui)
}

func TestFloat(t *testing.T) {
	fmt.Printf("%e", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%E", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%f", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%g", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%G", 12675757563.5345432567)
	fmt.Println()
}

func TestBool(t *testing.T) {
	fmt.Printf("%t", true)
	fmt.Println()
}

func TestString(t *testing.T) {
	str := "I'm a boy"
	fmt.Printf("%s", str)
	fmt.Println()
	fmt.Printf("%q", str)
	fmt.Println()
	fmt.Printf("%x", str)
	fmt.Println()
	fmt.Printf("%X", str)
	fmt.Println()
}

func TestPointer(t *testing.T) {
	a := "I'm a boy"
	b := &a
	fmt.Printf("%p", b)
	fmt.Println()
}

func TestPos(t *testing.T) {
	f := 1010.0101
	s := "hey boy!"
	fmt.Printf("%v", f)
	fmt.Println()

	fmt.Printf("%v", s)
	fmt.Println()
}
