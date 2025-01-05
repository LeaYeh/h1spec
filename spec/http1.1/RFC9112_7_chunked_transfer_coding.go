package http11

import "github.com/LeaYeh/h1spec/spec"

// HTTP11ChunkedTransferCoding implements tests for RFC9112 Section 7: "Chunked Transfer Coding".
func HTTP11ChunkedTransferCoding() *spec.TestGroup {
    tg := NewTestGroup("RFC9112.7", "Chunked Transfer Coding")
    // Add subchapter-level test groups and DO NOT implement the HTTP11ChunkSizeHexadecimal in this file
    // Reference Implementation will be in the future files
    tg.AddTestGroup(HTTP11ChunkSizeHexadecimal())
    return tg
}
