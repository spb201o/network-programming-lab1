package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
    conn, err := net.Dial("tcp", os.Args[1])

    if err != nil {
        os.Exit(-1)
    }

    for { 
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("> ")
        text, err := reader.ReadString('\n')
        if err != nil {
            os.Exit(-2)
        }

        fmt.Fprintf(conn, text + "\n")
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            os.Exit(-3)
        }

        fmt.Print(message)
    }
}
