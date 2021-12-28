package main

// Run `go test`
func TestName(t *testing.T) {
	name := getName()

	if name != "World!" {
		t.Error("Response from getName is unexpectedvalue")
	}
}
