package spec

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

var statusReasonMap = map[int]string{
	http.StatusContinue:           "Continue",
	http.StatusSwitchingProtocols: "Switching Protocols",
	http.StatusProcessing:         "Processing",
	http.StatusEarlyHints:         "Early Hints",

	http.StatusOK:                   "OK",
	http.StatusCreated:              "Created",
	http.StatusAccepted:             "Accepted",
	http.StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	http.StatusNoContent:            "No Content",
	http.StatusResetContent:         "Reset Content",
	http.StatusPartialContent:       "Partial Content",
	http.StatusMultiStatus:          "Multi-Status",
	http.StatusAlreadyReported:      "Already Reported",
	http.StatusIMUsed:               "IM Used",

	http.StatusMultipleChoices:   "Multiple Choices",
	http.StatusMovedPermanently:  "Moved Permanently",
	http.StatusFound:             "Found",
	http.StatusSeeOther:          "See Other",
	http.StatusNotModified:       "Not Modified",
	http.StatusUseProxy:          "Use Proxy",
	http.StatusTemporaryRedirect: "Temporary Redirect",
	http.StatusPermanentRedirect: "Permanent Redirect",

	http.StatusBadRequest:                  "Bad Request",
	http.StatusUnauthorized:                "Unauthorized",
	http.StatusPaymentRequired:             "Payment Required",
	http.StatusForbidden:                   "Forbidden",
	http.StatusNotFound:                    "Not Found",
	http.StatusMethodNotAllowed:            "Method Not Allowed",
	http.StatusNotAcceptable:               "Not Acceptable",
	http.StatusProxyAuthRequired:           "Proxy Authentication Required",
	http.StatusRequestTimeout:              "Request Timeout",
	http.StatusConflict:                    "Conflict",
	http.StatusGone:                        "Gone",
	http.StatusLengthRequired:              "Length Required",
	http.StatusPreconditionFailed:          "Precondition Failed",
	http.StatusRequestEntityTooLarge:       "Payload Too Large",
	http.StatusRequestURITooLong:           "URI Too Long",
	http.StatusUnsupportedMediaType:        "Unsupported Media Type",
	http.StatusRequestedRangeNotSatisfiable:"Range Not Satisfiable",
	http.StatusExpectationFailed:           "Expectation Failed",
	http.StatusTeapot:                      "I'm a teapot", // RFC 2324 humorous status code
	http.StatusMisdirectedRequest:          "Misdirected Request",
	http.StatusUnprocessableEntity:         "Unprocessable Entity",
	http.StatusLocked:                      "Locked",
	http.StatusFailedDependency:            "Failed Dependency",
	http.StatusTooEarly:                    "Too Early",
	http.StatusUpgradeRequired:             "Upgrade Required",
	http.StatusPreconditionRequired:        "Precondition Required",
	http.StatusTooManyRequests:             "Too Many Requests",
	http.StatusRequestHeaderFieldsTooLarge: "Request Header Fields Too Large",
	http.StatusUnavailableForLegalReasons:  "Unavailable For Legal Reasons",

	http.StatusInternalServerError:           "Internal Server Error",
	http.StatusNotImplemented:                "Not Implemented",
	http.StatusBadGateway:                    "Bad Gateway",
	http.StatusServiceUnavailable:            "Service Unavailable",
	http.StatusGatewayTimeout:                "Gateway Timeout",
	http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	http.StatusInsufficientStorage:           "Insufficient Storage",
	http.StatusLoopDetected:                  "Loop Detected",
	http.StatusNotExtended:                   "Not Extended",
	http.StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

var (
	gray   = color.New(color.FgHiBlack).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func DummyLongString(n int) string {
	var longString string
	for i := 0; i < n; i++ {
		longString += "a"
	}
	return longString
}

func StatusString(version float32, code int, end string) string {
	reason, ok := statusReasonMap[code]

	if !ok {
		return fmt.Sprintf("HTTP/%.1f %d Unknown Status Code%s", version, code, end)
	}
	return fmt.Sprintf("HTTP/%.1f %d %s%s", version, code, reason, end)
}

func FindInSlice(slice []string, str string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}
