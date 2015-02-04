// Implementing the basics of Go
package main

import "fmt"

type myString string

type pet struct {
	animal     string
	name       myString
	age        int
	knowsTrick bool
}

func main() {
	message1 := "Hello World!"
	var message2 myString = "Hello Mars!"
	var messagePtr *myString = &message2
	fmt.Println(message2, message1, *messagePtr)

	*messagePtr = "The average of:"
	x, y, z := 72, 54, 24
	fmt.Println(*messagePtr, x, y, z, "is", (x+y+z)/3)

	var pet1 = pet{"dog", "Fido", 3, true}
	fmt.Println("I have a", pet1.animal, "named", pet1.name, "who is", pet1.age, "years old.")
	if pet1.knowsTrick {
		fmt.Println(pet1.name, "is awesome so he knows tricks!")
	} else {
		fmt.Println(pet1.name, "is lazy so he doesn't know tricks.")
	}

	var pet2 pet
	pet2.animal = "cat"
	pet2.name = "Garfield"
	pet2.age = 5
	pet2.knowsTrick = false

	fmt.Println("I also have a", pet2.animal, "named", pet2.name, "who is", pet2.age, "years old.")
	if pet2.knowsTrick {
		fmt.Println(pet2.name, "is awesome so he knows tricks!")
	} else {
		fmt.Println(pet2.name, "is lazy so he doesn't know tricks.")
	}
}
