package day3

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings1(t *testing.T) {
	a := "hello"
	fmt.Println([]byte(a))
}

func TestStrings2(t *testing.T) {
	a := "hello"
	b := []byte(a)
	b[0] = 'x'
	fmt.Println(string(b))
}

func TestStrings3(t *testing.T) {
	a := "hello"
	fmt.Println(len(a), a[0], a[1:3])
}

func TestString4(t *testing.T) {
	fmt.Println(strings.Compare("ab", "cd"))
	fmt.Println(strings.EqualFold("ab", "AB"))
}

func TestString5(t *testing.T) {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(len("谷歌中国"), []byte("谷歌中国"))
	fmt.Println(strings.Count("谷歌中国", ""))

	for _, v := range []byte("谷歌中国") {
		fmt.Printf("%b\n", v)
	}
}
