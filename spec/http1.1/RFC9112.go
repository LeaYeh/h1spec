package http11

import "github.com/LeaYeh/h1spec/spec"

// RFC9112 is the main function for the RFC9112 protocol.
// It creates a new test group for the protocol and adds chapter-level test groups to it.
// The purpose of RFC9112 is to define the syntax and semantics of HTTP/1.1 request-target.
func RFC9112() *spec.TestGroup {
	tg := NewTestGroup("RFC9112", "Protocol RFC9112")

	// Add chapter-level test groups
	// The implementation of these test groups will be provided in future files
	tg.AddTestGroup(HTTP11RequestTarget())

	return tg
}
