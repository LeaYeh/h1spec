package http11

import "github.com/LeaYeh/h1spec/spec"

var key = "http1.1"

func NewTestGroup(section string, name string) *spec.TestGroup {
	return &spec.TestGroup{
		Key:     key,
		Section: section,
		Name:    name,
	}
}

func Spec() *spec.TestGroup {
	tg := &spec.TestGroup{
		Key:  key,
		Name: "Hypertext Transfer Protocol Version 1.1 (HTTP/1.1)",
	}

	// Add the main chapters as test groups
	tg.AddTestGroup(RFC7230())
	// tg.AddTestGroup(RFC7231())
	// tg.AddTestGroup(RFC7232())
	// tg.AddTestGroup(RFC7233())
	// tg.AddTestGroup(RFC7234())
	// tg.AddTestGroup(RFC7235())

	return tg
}
