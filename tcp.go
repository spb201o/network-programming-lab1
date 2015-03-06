package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:1145")

    if err != nil {
        os.Exit(-11)
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