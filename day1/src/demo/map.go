package main

import "fmt"

func main() {
	var m1 map[string]string // map[int]string  {"key1": "vaule1"}
	fmt.Println(m1)          // {}

	var m2 = map[string]string{"key1": "value1"}
	m2["key2"] = "value2"
	fmt.Println(m2) //{"key1": "value1", "key2": "value2"}
	m2["key2"] = "value2_edit"
	fmt.Println(m2)

	fmt.Println(m2)
	fmt.Println(len(m2))

	//
	for k, v := range m2 {
		fmt.Println(k, v)
	}
}
