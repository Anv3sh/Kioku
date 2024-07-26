package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    serverAddr := "localhost"
    serverPort := "6379"

    // Establish a TCP connection to the Telnet server.
    conn, err := net.Dial("tcp", serverAddr+":"+serverPort)
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        return
    }
    defer conn.Close()

    // reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)

    // Interactive loop to send commands and receive responses.
    // go func() {
    //     for {
    //         buf := make([]byte, 1024)
    //         // Read server response and print it.
    //         _, err:=conn.Read(buf)
    //         if err !=nil{
    //             fmt.Printf("read data failed: ",buf)
    //         }
    //         fmt.Printf(string(buf))
    //     }
    // }()

    // Read user input and send it to the server.
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        command := scanner.Text()
        writer.WriteString(command + "\r\n") // Telnet requires \r\n for line endings
        writer.Flush()
    }
}
