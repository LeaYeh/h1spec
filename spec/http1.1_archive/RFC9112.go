package http11

import "github.com/LeaYeh/h1spec/spec"

func RFC9112() *spec.TestGroup {
	tg := NewTestGroup("RFC9112", "Protocol RFC9112")

	// IMPORTANT: Preserve existing test groups
    tg.AddTestGroup(HTTP11RequestTarget())

	// Add chapter-level test group
	// RFC9112.7: Chunked Transfer Coding
	// This test group is to verify the implementation of chunked transfer coding in HTTP/1.1
	// Reference Implementation will be in the future files
	tg.AddTestGroup(HTTP11ChunkedTransferCoding())

	return tg
}
