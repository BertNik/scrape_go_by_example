




package main


import "fmt"


func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}


func main() {


    nextInt := intSeq()


    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())


$ go run closures.go
1
2
3
1

