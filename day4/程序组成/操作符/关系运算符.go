package main

import "fmt"

func main() {
	fmt.Printf("%T, %T \n", a, b)
	// fmt.Println(a == b)
}

func a(a int) {
	fmt.Println(a)
}

func b(a int) {
	fmt.Println(a)
}
