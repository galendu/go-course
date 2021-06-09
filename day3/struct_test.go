package day3

import (
	"fmt"
	"testing"
)

type Person struct {
	Name          string
	Age           int
	Gender        string
	Weight        uint
	FavoriteColor []string
}

func TestStructDel1(t *testing.T) {
	var person Person
	fmt.Printf("%+v\n", person)
}

func TestStructDel2(t *testing.T) {
	var person Person = Person{
		Name:          "andy",
		Age:           66,
		Gender:        "male",
		Weight:        120,
		FavoriteColor: []string{"red", "blue"},
	}
	fmt.Printf("%+v\n", person)
}

func TestStructP1(t *testing.T) {
	var person *Person
	fmt.Println(person)
}

func TestStructP2(t *testing.T) {
	var person *Person = &Person{
		Name:          "andy",
		Age:           66,
		Gender:        "male",
		Weight:        120,
		FavoriteColor: []string{"red", "blue"},
	}
	fmt.Printf("%p", person)
}

func TestStructP3(t *testing.T) {
	person := new(Person)
	fmt.Printf("%p", person)
}
