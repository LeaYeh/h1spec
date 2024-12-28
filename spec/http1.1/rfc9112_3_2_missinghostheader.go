package http11

import (
    "github.com/LeaYeh/h1spec/config"
    "github.com/LeaYeh/h1spec/spec"
)

// HTTP11MissingHostHeader implements tests for RFC9112 Section 3.2: "Request Target".
func HTTP11MissingHostHeader() *spec.TestGroup {
    tg := NewTestGroup("RFC9112.3.2", "Request Target")
    
    tg.AddTestCase(&spec.TestCase{
        Strict:      true, // true or false, based on the Mode
        Desc:        "A server MUST respond with a 400 (Bad Request) status code to any HTTP/1.1 request message that lacks a Host header field and to any request message that contains more than one Host header field line or a Host header field with an invalid field value.",
        Requirement: "The server MUST respond with a 400 (Bad Request) status code to any HTTP/1.1 request message that lacks a Host header field and to any request message that contains more than one Host header field line or a Host header field with an invalid field value.",
        Run: func(c *config.Config, conn *spec.Conn) error {
            expected := []string{spec.StatusString(1.1, 400, "\r")}
            request := "GET / HTTP/1.1\r\n" +
                    "Connection: close\r\n\r\n"

            err := conn.Send([]byte(request))
            if err != nil {
                return err
            }
            acturl, err := conn.ReadLine()
            if err != nil {
                return err
            }

            if !spec.FindInSlice(expected, acturl) {
                return &spec.TestError{
                    Expected: expected,
                    Actual:   acturl,
                }
            }
            return nil
        },
    })
    
    return tg
}
