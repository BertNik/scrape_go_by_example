




package main


import "crypto/sha1"
import "fmt"


func main() {
    s := "sha1 this string"


    h := sha1.New()


    h.Write([]byte(s))


    bs := h.Sum(nil)


$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94



    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}


