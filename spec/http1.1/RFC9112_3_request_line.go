// spec/http1.1/RFC9112_3_request_line.go
package http11

import (
	"github.com/LeaYeh/h1spec/spec"
    "github.com/LeaYeh/h1spec/config"
)

// ### ✅ TC-RFC9112-3.0-RequestLine-Valid-P0

// * **Test Description**: Ensure that the server accepts a syntactically valid request-line.
// * **Preconditions**: Server is running and listening on a TCP socket.
// * **Test Steps**:

//   1. Send: `GET /index.html HTTP/1.1\r\n`
//   2. Include a valid `Host` header.
// * **Expected Result**: `HTTP/1.1 200 OK` or other valid response.
// * **Actual Result**: *(To be filled)*
// * **Status**: *(Pass / Fail / In Progress)*
// * **Comments/Notes**: Baseline validation for correctly formatted requests.

// HTTP11RequestLine implements tests for RFC9112 Section 3: "Request Line".
func HTTP11RequestLine() *spec.TestGroup {
    tg := NewTestGroup("RFC9112.3", "Request Line")

    tg.AddTestCase(&spec.TestCase{
        Strict:      false,
        Key:         "TC-RFC9112-3.0-RequestLine-Valid-P0",
        Desc:        "Ensure that the server accepts a syntactically valid request-line",
        Requirement: "The request-line MUST be in the format: `Method SP Request-URI SP HTTP-Version CRLF`.",
        Run: func(c *config.Config, conn *spec.Conn) error {
            expected := []string{spec.StatusString(1.1, 200, "\r")}
            request := "GET /index.html HTTP/1.1\r\n" +
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

    // Add subchapter-level test groups
    return tg
}
