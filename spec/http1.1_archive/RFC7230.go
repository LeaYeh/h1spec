package http11

import "github.com/LeaYeh/h1spec/spec"

// RFC7230 is the main function for the RFC7230 protocol level.
// It creates a new test group for the RFC7230 and adds the existing and new chapter-level test groups to it.
// The purpose of this function is to test the HTTP/1.1 protocol as defined in the RFC7230.
func RFC7230() *spec.TestGroup {
	tg := NewTestGroup("RFC7230", "Protocol RFC7230")

	// IMPORTANT: Preserve existing test groups
	tg.AddTestGroup(HTTP11Architecture())

	// Add chapter-level test group for Message Body
	// This test group is for testing the rules for constructing HTTP/1.1 message bodies as defined in RFC7230 Section 3.3.3
	// Reference Implementation will be in the future files
	tg.AddTestGroup(HTTP11MessageBody())

	return tg
}
