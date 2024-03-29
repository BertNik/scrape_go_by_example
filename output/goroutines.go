




package main


import "fmt"


func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}


func main() {


    f("direct")


    go f("goroutine")


    go func(msg string) {
        fmt.Println(msg)
    }("going")


$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done


    fmt.Scanln()
    fmt.Println("done")
}


