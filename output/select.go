




package main


import "time"
import "fmt"


func main() {


    c1 := make(chan string)
    c2 := make(chan string)


    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()


$ time go run select.go 
received one
received two

