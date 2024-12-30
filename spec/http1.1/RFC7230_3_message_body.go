package http11

import "github.com/LeaYeh/h1spec/spec"

// HTTP11MessageBody implements tests for RFC7230 Section 3: "Message Body".
func HTTP11MessageBody() *spec.TestGroup {
    tg := NewTestGroup("RFC7230.3", "Message Body")
    // Add subchapter-level test groups and DO NOT implement the HTTP11MultipleContentLength in this file
    // Reference Implementation will be in the future files
    tg.AddTestGroup(HTTP11MultipleContentLength())
    return tg
}
