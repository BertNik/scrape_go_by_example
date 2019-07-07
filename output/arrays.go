




package main


import "fmt"


func main() {


    var a [5]int
    fmt.Println("emp:", a)


    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])


    fmt.Println("len:", len(a))


    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

/*
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]
*/

