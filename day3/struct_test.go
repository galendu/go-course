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
	NewAttr       string
	Addr          Home
}

type Home struct {
	City string
	T1   T1
}

type T1 struct {
	T1 string
}

func (p Person) Add() int {
	return p.Age * 2
}

type Author struct {
	Name string
	Aage int
}

func (a *Author) GetName() string {
	return a.Name
}

type Titile struct {
	Main string
	Sub  string
}

type Book struct {
	Author *Author
	Titile *Titile
}

func (b *Book) GetName() string {
	return b.Author.GetName() + "book"
}

func TestMain(t *testing.T) {
	b1 := Book{
		Author: &Author{
			Name: "laoyu",
		},
		Titile: &Titile{},
	}

	b2 := &Book{
		Author: &Author{},
		Titile: &Titile{},
	}
	*b2.Author = *b1.Author
	*b2.Titile = *b1.Titile

	b2.Author.Name = "new author"

	fmt.Println(b1.Author.Name, b2.Author.Name)

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
