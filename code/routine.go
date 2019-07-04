package main

import "fmt"

func main() {

    // Create a new channel with
    messages := make(chan string)

    // Send a value into a channel using the '<-'
    go func() { messages <- "ping" }()

    // Here we'll receive the "ping" message
    // we sent above and print it out.
    msg := <-messages
    fmt.Println(msg)
}
