




package main


import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func main() {


    scanner := bufio.NewScanner(os.Stdin)


    for scanner.Scan() {


        ucl := strings.ToUpper(scanner.Text())


        fmt.Println(ucl)
    }


$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

