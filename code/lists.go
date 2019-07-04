package main

import "fmt"

func main() {
	var arr [3]int                        //initiate array [0.0.0]
	arr[0] = 4                            //assign value to element
	c := []string{"blue", "black", "red"} //initiate SLICE - when you don't know array size
	b := make([]int, 0, 5)                // len(b)=0, cap(b)=5

	c = append(c, "orange") //append to slice

	fmt.Printf("len=%d cap=%d c=%v\n", len(c), cap(c), c)
	fmt.Printf("len=%d cap=%d b=%v\n", len(b), cap(b), b)
}
