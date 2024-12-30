package http11

import "github.com/LeaYeh/h1spec/spec"

// RFC7230 is the main function for the RFC7230 test group.
// It includes all the chapter-level test groups for RFC7230.
func RFC7230() *spec.TestGroup {
	tg := NewTestGroup("RFC7230", "Protocol RFC7230")

	// IMPORTANT: Preserve existing test groups
	tg.AddTestGroup(HTTP11Architecture())

	// Add chapter-level test group for Message Body
	// This test group will test the rules for the message body as defined in RFC7230 Section 3.3.3
	// Reference Implementation will be in the future files
	tg.AddTestGroup(HTTP11MessageBody())

	return tg
}
