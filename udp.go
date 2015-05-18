package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
    serverAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:1145")
    conn, err := net.DialUDP("udp", nil, serverAddr)

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
        byteArray := make([]byte, 1500)
        n, _, err := conn.ReadFromUDP(byteArray)
        if err != nil {
            os.Exit(-3)
        }

        fmt.Println(string(byteArray[:n]))
    }
}
