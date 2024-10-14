package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
    conn, err := net.Dial("tcp", "webserv_container:4000")
    // conn, err := net.Dial("tcp", "webserv_container:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    requestLine := "GET / HTTP/1.1\r\n"
    headers := "" +
                "Host: webserv_container\r\n" +
                "Connection: Close\r\n" +
                "\r\n"

    // requestLine := "POST /upload/file2 HTTP/1.1\r\n"
    // headers := "Host: webserv_container\r\n" +
    //            "Connection: Close\r\n" +
    //            "Content-Type: text/plain\r\n" +
    //            "Content-Length: %d\r\n" + // Placeholder for the body length
    //             "\r\n"

    // body := "This is the body of the request"
    // headers = fmt.Sprintf(headers, len(body)) + body + "\r\n"

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
