package http11

import "github.com/LeaYeh/h1spec/spec"

// HTTP11RequestTarget implements tests for RFC9112 Section 3: "Request Target".
func HTTP11RequestTarget() *spec.TestGroup {
    tg := NewTestGroup("RFC9112.3", "Request Line")
    // Add subchapter-level test groups
    tg.AddTestGroup(HTTP11MissingHostHeader())
    return tg
}
