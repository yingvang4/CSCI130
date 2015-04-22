package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

type myString string

func (p Person) PrintInfo() {
	fmt.Println(p.name, "is ", p.age, "years old!")
}

func (s myString) AddsAwesomeness() (Awesomeness myString) {
	Awesomeness = s + " Awesomeness"
	return
}

func main() {
	SliceName := []string{"Bobby","Tom","Jerry"}
	Winner := []string{"GoodGuy"}
	
	fmt.Println("List of civilians:", SliceName)
	fmt.Println("")
	fmt.Println("EvilGuy enters the scene and kills Jerry")
	SliceName = append(SliceName[:2], "EvilGuy")
	fmt.Println(SliceName)
	fmt.Println("")
	fmt.Println("Civilian leaves the scene as GoodGuy enters the scene")
	SliceName = append(SliceName[2:],"GoodGuy")
	fmt.Println(SliceName)
	fmt.Println("")
	
	if SliceName[1] == Winner[0] {
		SliceName = SliceName[1:]
		fmt.Println(SliceName ," wins")
	}
	fmt.Println("")
	
	p1 := Person{"Joseph",20}
	p2 := Person{"Tyler", 14}
	
	p1.PrintInfo()
	p2.PrintInfo()
	
	var x myString = "This is"
	fmt.Println(x.AddsAwesomeness())
	
	
}
