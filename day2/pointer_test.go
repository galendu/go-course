package day2

import (
	"fmt"
	"testing"
)

func TestPointerRaw(t *testing.T) {
	a := "string a"
	fmt.Println(&a)
}
