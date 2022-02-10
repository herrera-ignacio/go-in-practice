package template

// go test -bench . -cpu=1,2,4

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkParallelTemplates(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}

	b.RunParallel(func(pb *testing.PB) {
		// Instead of a for loop, passes a closure into RunParallel
		var buf bytes.Buffer
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
