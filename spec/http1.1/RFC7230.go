package http11

import "github.com/LeaYeh/h1spec/spec"

func RFC7230() *spec.TestGroup {
	tg := NewTestGroup("RFC7230", "Protocol RFC7230")

	tg.AddTestGroup(HTTP11Architecture())

	return tg
}