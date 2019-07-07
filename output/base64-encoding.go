




package main


import b64 "encoding/base64"
import "fmt"


func main() {


    data := "abc123!?$*&()'-=@~"


    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)


    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()


$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

