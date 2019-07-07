




package main


import "fmt"


func main() {


    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)


$ go run range-over-channels.go
one
two


    for elem := range queue {
        fmt.Println(elem)
    }
}


