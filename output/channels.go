




package main


import "fmt"


func main() {


    messages := make(chan string)


    go func() { messages <- "ping" }()


$ go run channels.go 
ping


    msg := <-messages
    fmt.Println(msg)
}


