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

	return tg
}
