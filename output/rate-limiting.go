




package main


import "time"
import "fmt"


func main() {


    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)


    limiter := time.Tick(200 * time.Millisecond)


    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }


    burstyLimiter := make(chan time.Time, 3)


    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }


    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()


$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

