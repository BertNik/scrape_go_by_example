




package main


import "syscall"
import "os"
import "os/exec"


func main() {


    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }


    args := []string{"ls", "-a", "-l", "-h"}


    env := os.Environ()


$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

