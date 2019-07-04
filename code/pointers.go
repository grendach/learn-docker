package main

import "fmt"

func main() {
	name := "Dmytro"
	age := 33
	p := &name // point to name
	p1 := &age
	fmt.Printf("Hello, my name is %s and my memory location is: %p\n", name, p)
	fmt.Printf("%s is %d years old\n", *p, *p1)
	*p1 = *p1 + 100 // assign to "age" new value through it's pointer
	// this is known as "dereferencing" or "indirecting"
	fmt.Printf("%s is goint to live till he will be %d years old\n", *p, *p1)
}
