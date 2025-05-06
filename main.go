package main

import "fmt"

type Speaker interface {
	Speak() string
}
type Cat struct {
}

type Dog struct {
}

type Person struct {
	name string
}

func main() {
	c := Cat{}
	d := Dog{}
	p := Person{name: "John"}

	saySomething(c, d, p)
}

func (c Cat) Speak() string {
	return "Meow"
}

func (d Dog) Speak() string {
	return "Woof"
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hello, my name is %s", p.name)
}

func saySomething(s ...Speaker) {
	for _, speaker := range s {
		fmt.Println(speaker.Speak())
	}
}
