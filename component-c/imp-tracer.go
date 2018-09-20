package main


import (
    "fmt"
    "os"
)

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Need only one argument")
        os.Exit(1)
    }

    script := args[1]
    cmd := exec.Command(script)

    err := cmd.Run()

    if err != nil {
        fmt.Println("Step execution is successfull!")
    } else {
        fmt.Println("Step execution failed")
    }
}
