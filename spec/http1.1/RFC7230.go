package http11

import "github.com/LeaYeh/h1spec/spec"

// RFC7230 is the main function for the RFC7230 protocol level.
// It creates a new test group for the RFC7230 protocol and adds chapter-level test groups to it.
// The RFC7230 protocol defines the syntax and semantics of HTTP/1.1 messages.
// The chapter-level test groups are named as HTTP11{{CHAPTER_TITLE}}, where CHAPTER_TITLE is the title of the chapter.
func RFC7230() *spec.TestGroup {
	tg := NewTestGroup("RFC7230", "Protocol RFC7230")

	// Add chapter-level test groups
	// The implementation of these test groups will be in future files
	tg.AddTestGroup(HTTP11MessageFormat())
	tg.AddTestGroup(HTTP11ConnectionManagement())
	tg.AddTestGroup(HTTP11MessageSemantics())
	tg.AddTestGroup(HTTP11PayloadSemantics())
	tg.AddTestGroup(HTTP11ConditionalRequests())
	tg.AddTestGroup(HTTP11RangeRequests())
	tg.AddTestGroup(HTTP11Caching())
	tg.AddTestGroup(HTTP11Authentication())

	return tg
}
