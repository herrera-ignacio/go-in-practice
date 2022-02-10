package main

/**
go test -bench . --race -cpu=1,2,4
*/

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkParallelOops(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}

	var buf bytes.Buffer // moved out of the closure
	/**
	It's likely that the parallel test will fail. If you add the `-race`
	flag, Go instruments for race conditions and you receive more helpful
	information.
	The reason is that `bytes.Buffer` is tried to be used by more than one
	thing at once.
	From here you have the options of stepping back to declaring the buffer
	inside the closure, or using a `sync.Mutex` to lock and unlock access.
	*/
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
