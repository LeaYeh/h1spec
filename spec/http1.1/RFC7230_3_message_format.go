package http11

import "github.com/LeaYeh/h1spec/spec"

// HTTP11MessageFormat implements tests for RFC7230 Section 3: "Message Format".
func HTTP11MessageFormat() *spec.TestGroup {
    tg := NewTestGroup("RFC7230.3", "Message Format")
    // Add subchapter-level test groups
    tg.AddTestGroup(HTTP11MultipleContentLength())
    return tg
}