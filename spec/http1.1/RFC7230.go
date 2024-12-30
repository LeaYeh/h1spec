package http11

import "github.com/LeaYeh/h1spec/spec"

// RFC7230 is the main function for the RFC7230 protocol.
// It creates a new test group for the protocol and adds chapter-level test groups to it.
// The function returns a pointer to the created test group.
func RFC7230() *spec.TestGroup {
	tg := NewTestGroup("RFC7230", "Protocol RFC7230")

	// Add chapter-level test group for Message Format
	// The actual implementation of the test cases will be in the future files
	tg.AddTestGroup(HTTP11MessageFormat())

	return tg
}
