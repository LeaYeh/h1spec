package http11

import (
	"fmt"

	"github.com/LeaYeh/h1spec/config"
	"github.com/LeaYeh/h1spec/spec"
)

func Http11ErrorHandling() *spec.TestGroup {
	tg := NewTestGroup("RFC7230.2.5", "Conformance and Error Handling")

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The request URI length is out of the server's capability, need to return 414",
		Requirement: "The recipient MUST be able to parse any value of reasonable length that is applicable.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			expected := []string{spec.StatusString(1.1, 414, "\r")}
			request := "GET /" + spec.DummyLongString(100000) + " HTTP/1.1\r\n" +
					"Host: " + c.Host + "\r\n" +
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

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The request header length is out of the server's capability, need to return 431 or 400",
		Requirement: "The recipient MUST be able to parse any value of reasonable length that is applicable.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			expected := []string{spec.StatusString(1.1, 414, "\r"), spec.StatusString(1.1, 400, "\r")}

			request := "GET / HTTP/1.1\r\n" +
					"Host: " + c.Host + "\r\n" +
					"Connection: close\r\n" +
					"Cookie: " + spec.DummyLongString(10000) + "\r\n\r\n"

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

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "Test for body exceeding server limit, expecting 413 Payload Too Large",
		Requirement: "The server must reject requests with a body that exceeds its size limits.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			expected := []string{spec.StatusString(1.1, 413, "\r")}

			body := spec.DummyLongString(100000)
			request := "POST /upload/file2 HTTP/1.1\r\n" +
				"Host: " + c.Host + "\r\n" +
				"Content-Length: " + fmt.Sprint(len(body)) + "\r\n" +
				"Connection: close\r\n\r\n" +
				body + "\r\n"

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