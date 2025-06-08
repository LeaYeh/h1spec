package http11

import (
    "github.com/LeaYeh/h1spec/config"
    "github.com/LeaYeh/h1spec/spec"
)

// HTTP11MultipleContentLength implements tests for RFC7230 Section 3.3.3: "Message Body Length".
func HTTP11MultipleContentLength() *spec.TestGroup {
    tg := NewTestGroup("RFC7230.3.3.3", "Message Body Length")
    
    tg.AddTestCase(&spec.TestCase{
        Strict:      true, // true or false, based on the Mode
        Desc:        "Multiple Content-Length values in a request",
        Requirement: "If a message is received without Transfer-Encoding and with multiple Content-Length header fields having differing field-values, then the message framing is invalid and the recipient MUST treat it as an unrecoverable error.  If this is a request message, the server MUST respond with a 400 (Bad Request) status code and then close the connection.",
        Run: func(c *config.Config, conn *spec.Conn) error {
            expected := []string{spec.StatusString(1.1, 400, "\r")}
            request := "POST /submit HTTP/1.1\r\n" +
                "Host: www.hehe.com\r\n" +
                "Content-Length: 2000\r\n" +
                "Content-Length: 3\r\n\r\n"

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
