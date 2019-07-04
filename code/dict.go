package main

import "fmt"

func main() {
	//initianilisation
	go_map := make(map[string]int)
	
	//assign key, value
	go_map["my_key"] = 5
	fmt.Println(go_map["my_key"])
	
	// change value
	go_map["my_key"] = 555
	fmt.Println(go_map["my_key"])
	
	//remove by key
	delete(go_map, "my_key")
	fmt.Println(go_map)
}
