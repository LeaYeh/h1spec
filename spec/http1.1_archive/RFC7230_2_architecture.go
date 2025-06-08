package http11

import "github.com/LeaYeh/h1spec/spec"

func HTTP11Architecture() *spec.TestGroup {
	tg := NewTestGroup("RFC7230.2", "Architecture")

	tg.AddTestGroup(Http11ErrorHandling())
	tg.AddTestGroup(Http11ProtocolVersioning())

	return tg
}