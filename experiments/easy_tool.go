package main

import (
    "fmt"
    "net"
    "bufio"
)

func main() {
    conn, err := net.Dial("tcp", "iana.org:443")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    requestLine := "GET / HTTP/1.1\r\n"
    headers := ""
    // headers += "Host: iana.org\r\n\r\n"
    headers += "Connection: close\r\n\r\n"

    httpRequest := requestLine + headers
    fmt.Println(httpRequest)

    fmt.Fprintf(conn, httpRequest)

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading response:", err)
    }
}
