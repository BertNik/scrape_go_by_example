




package main


import "os"


func main() {


    panic("a problem")


$ go run panic.go
panic: a problem


goroutine 1 [running]:
main.main()
    /.../panic.go:12 +0x47
...
exit status 2

