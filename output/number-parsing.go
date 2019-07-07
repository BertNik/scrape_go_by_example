




package main


import "strconv"
import "fmt"


func main() {


    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)


    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)


    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)


    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)


    k, _ := strconv.Atoi("135")
    fmt.Println(k)


$ go run number-parsing.go 
1.234
123
456
789
135
strconv.ParseInt: parsing "wat": invalid syntax
