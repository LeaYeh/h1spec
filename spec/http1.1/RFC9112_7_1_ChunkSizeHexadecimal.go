package http11

import (
    "github.com/LeaYeh/h1spec/config"
    "github.com/LeaYeh/h1spec/spec"
)

// HTTP11ChunkSizeHexadecimal implements tests for RFC9112 Section 7.1: "Chunked Transfer Coding".
func HTTP11ChunkSizeHexadecimal() *spec.TestGroup {
    tg := NewTestGroup("RFC9112.7.1", "Chunked Transfer Coding")
    
    tg.AddTestCase(&spec.TestCase{
        Strict:      true, // true or false, based on the Mode
        Desc:        "Chunk Size can only contain hexadecimal characters (0123456789ABCDEFabcdef)",
        Requirement: "The chunk-size field is a string of characters, each character being a hexadecimal digit (case-insensitive).",
        Run: func(c *config.Config, conn *spec.Conn) error {
            expected := []string{spec.StatusString(1.1, 400, "\r")}
            request := "POST / HTTP/1.1\r\n" +
                "Host: www.example.com\r\n" +
                "Transfer-Encoding: chunked\r\n\r\n" +
                "A\r\n" +
                "0123456789\r\n" +
                "1\r\n" +
                "a\r\n" +
                "2\r\n" +
                "Hi\r\n" +
                "+a\r\n" +
                "0123456789\r\n" +
                "0\r\n\r\n"

            err := conn.Send([]byte(request))
            if err != nil {
                return err
            }
            actual, err := conn.ReadLine()
            if err != nil {
                return err
            }

            if !spec.FindInSlice(expected, actual) {
                return &spec.TestError{
                    Expected: expected,
                    Actual:   actual,
                }
            }
            return nil
        },
    })
    
    return tg
}
