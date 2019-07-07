




package main


import "time"
import "fmt"


func main() {


    timer1 := time.NewTimer(2 * time.Second)


    <-timer1.C
    fmt.Println("Timer 1 expired")

