package http11

import (
	"github.com/LeaYeh/h1spec/config"
	"github.com/LeaYeh/h1spec/spec"
)

func Http11ProtocolVersioning() *spec.TestGroup {
	tg := NewTestGroup("RFC7230.2.6", "HTTP/1.1 Protocol Versioning")

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The HTTP-version need to follow the format, HTTP/1 is invalid",
		Requirement: " HTTP-version  = HTTP-name \"/\" DIGIT \".\" DIGIT",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 400 Bad Request\r"

			request := "GET / HTTP/1\r\n" +
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

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The HTTP-name is case-insensitive, http is not a valid HTTP-name",
		Requirement: "HTTP-name     = %x48.54.54.50 ; \"HTTP\", case-sensitive",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 400 Bad Request\r"

			request := "GET / http/1.1\r\n" +
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

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "Assuming the server supports HTTP/1.1 only, request for HTTP/1.1 should work normally",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 200 OK\r"

			request := "GET / HTTP/1.1\r\n" +
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

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	// TODO: HTTP/0.9 is hard to disable in modern servers, maybe we can skip this test
	// tg.AddTestCase(&spec.TestCase{
	// 	Strict:      true,
	// 	Desc:        "Assuming the server supports HTTP/1.1 only, request for HTTP/0.9 should behave as if it was HTTP/1.1",
	// 	Requirement: "A recipient can assume that a message with a higher minor version, when sent to a recipient that has not yet indicated support for that higher version, is sufficiently backwards-compatible to be safely processed by any implementation of the same major version.",
	// 	Run: func(c *config.Config, conn *spec.Conn) error {
	// 		passed := true
	// 		expected := "HTTP/1.1 505 HTTP Version Not Supported\r"

	// 		request := "GET /\r\n"

	// 		err := conn.Send([]byte(request))
	// 		if err != nil {
	// 			return err
	// 		}
	// 		acturl, err := conn.ReadLine()
	// 		if err != nil {
	// 			return err
	// 		}

	// 		if acturl != expected {
	// 			passed = false
	// 		}

	// 		if !passed {
	// 			return &spec.TestError{
	// 				Expected: []string{expected},
	// 				Actual:   acturl,
	// 			}
	// 		}
	// 		return nil
	// 	},
	// })

	tg.AddTestCase(&spec.TestCase{
		Strict:      true,
		Desc:        "Assuming the server supports HTTP/1.1 only, request for HTTP/1.2 should behave as if it was HTTP/1.1",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 200 OK\r"

			request := "GET / HTTP/1.2\r\n" +
						  "Host: " + c.Host + "\r\n" +
						  "User-Agent: " + c.AgentName + "\r\n" +
						  "Connection: close\r\n\r\n"

			err := conn.Send([]byte(request))
			if err != nil {
				return err
			}
			acturl, err := conn.ReadLine()
			if err != nil {
				return err
			}

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "Assuming the server supports HTTP/1.1 only, request for HTTP/2.0 should behave as if it was HTTP/1.1",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 505 HTTP Version Not Supported\r"

			request := "GET / HTTP/2.0\r\n" +
                   "Host: " + c.Host + "\r\n" +
                   "User-Agent: " + c.AgentName + "\r\n" +
                   "Connection: close\r\n\r\n"

			err := conn.Send([]byte(request))
			if err != nil {
				return err
			}
			acturl, err := conn.ReadLine()
			if err != nil {
				return err
			}

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "Assuming the server supports HTTP/1.1 only, request for HTTP/3.0 should behave as if it was HTTP/1.1",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 505 HTTP Version Not Supported\r"

			request := "GET / HTTP/3.0\r\n" +
                   "Host: " + c.Host + "\r\n" +
                   "User-Agent: " + c.AgentName + "\r\n" +
                   "Connection: close\r\n\r\n"

			err := conn.Send([]byte(request))
			if err != nil {
				return err
			}
			acturl, err := conn.ReadLine()
			if err != nil {
				return err
			}

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The request for HTTP/-1.1 is invalid",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 400 Bad Request\r"

			request := "GET / HTTP/-1.1\r\n"
			err := conn.Send([]byte(request))
			if err != nil {
				return err
			}
			acturl, err := conn.ReadLine()
			if err != nil {
				return err
			}

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})


	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The request for HTTP/0.0 is invalid",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 400 Bad Request\r"

			request := "GET / HTTP/0.0\r\n"
			err := conn.Send([]byte(request))
			if err != nil {
				return err
			}
			acturl, err := conn.ReadLine()
			if err != nil {
				return err
			}

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The request for HTTP/0.2 is invalid",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 400 Bad Request\r"

			request := "GET / HTTP/0.2\r\n"
			err := conn.Send([]byte(request))
			if err != nil {
				return err
			}
			acturl, err := conn.ReadLine()
			if err != nil {
				return err
			}

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	tg.AddTestCase(&spec.TestCase{
		Strict:      false,
		Desc:        "The request for HTTP/4.2 is invalid",
		Requirement: "A server MUST NOT send a version to which it is not conformant. A server can send a 505.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			passed := true
			expected := "HTTP/1.1 505 HTTP Version Not Supported\r"

			request := "GET / HTTP/4.2\r\n\r\n"

			err := conn.Send([]byte(request))
			if err != nil {
				return err
			}
			acturl, err := conn.ReadLine()
			if err != nil {
				return err
			}

			if acturl != expected {
				passed = false
			}

			if !passed {
				return &spec.TestError{
					Expected: []string{expected},
					Actual:   acturl,
				}
			}
			return nil
		},
	})

	return tg
}

