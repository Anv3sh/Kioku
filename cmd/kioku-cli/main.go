package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Get server address and port from command-line arguments or user input.
    serverAddr := "localhost"
    serverPort := "6379" // Default Telnet port

    // Establish a TCP connection to the Telnet server.
    conn, err := net.Dial("tcp", serverAddr+":"+serverPort)
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        return
    }
    defer conn.Close()

    // Create reader and writer for the connection.
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)

    // Interactive loop to send commands and receive responses.
    go func() {
        for {
            // Read server response and print it.
            msg, _ := reader.ReadByte()
            fmt.Print(msg)
        }
    }()

    // Read user input and send it to the server.
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        command := scanner.Text()
        writer.WriteString(command + "\r\n") // Telnet requires \r\n for line endings
        writer.Flush()
    }
}
