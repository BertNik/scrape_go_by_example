




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


    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }
}


$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...

