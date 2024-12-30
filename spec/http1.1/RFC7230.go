package http11

import "github.com/LeaYeh/h1spec/spec"

// RFC7230 is the main function for the RFC7230 protocol.
// It creates a new test group for the protocol and adds chapter-level test groups to it.
// The function returns the created test group.
func RFC7230() *spec.TestGroup {
	tg := NewTestGroup("RFC7230", "Protocol RFC7230")

	// Add chapter-level test group for Message Body
	// The Message Body in HTTP messages is used to carry the payload body associated with a request or response.
	// This test group will contain test cases related to the rules and requirements of the Message Body as defined in RFC7230 Section 3.3.3.
	tg.AddTestGroup(HTTP11MessageBody())

	return tg
}
