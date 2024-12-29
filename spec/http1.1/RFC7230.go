package http11

import "github.com/LeaYeh/h1spec/spec"

// RFC7230 is the main function for the RFC7230 protocol.
// It creates a new test group for the protocol and adds chapter-level test groups to it.
// The purpose of this RFC is to define the syntax and semantics of HTTP/1.1 messages, both as they are sent on the wire and as they are understood in terms of HTTP semantics.
func RFC7230() *spec.TestGroup {
	tg := NewTestGroup("RFC7230", "Protocol RFC7230")

	// Add chapter-level test group for Message Format
	// The implementation of this test group will be done in future files
	tg.AddTestGroup(HTTP11MessageFormat())

	return tg
}
