




package main


import "os"
import "strings"
import "fmt"


func main() {


    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))


$ go run environment-variables.go
FOO: 1
BAR: 


TERM_PROGRAM
PATH
SHELL
...

