# Verifying interfaces with canary tests

- [Verifying interfaces with canary tests](#verifying-interfaces-with-canary-tests)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)
  - [Example](#example)

## Problem

You want to make sure that the interfaces you're defining describe the things that you're intending to describe. This is useful in four cases:

* When you're exporting types that implement external interfaces.
* When you create interface types that describe external types.
* When you rely on external interfaces, and those interfaces change even though they shouldn't.
* When use of that interface is restricted to type assertions.

> **Interfaces change infrequently**: Ideally, after an interface is exported and made public, it shouldn't be changed.

## Solution

Write type-assertion "canary" tests that will fail quickly if you made a mistake on your interface definition.

## Discussion

When you're writing interfaces, or implementations of interfaces, it's useful to write simple type-assertion canary tests that will explode at compile time.

> A canary test is a test designed to alert you of basic failures in your assumptions.

## Example

Say you're writing a customer writer that implements `io.Writer`. You're exporting this in your library so that other code may use it:

```go
type MyWriter struct {
  // ...
}

func (m *MyWriter) Write(p []byte) error {
  // Write data somewhere...
  return nil
}

func main() {
  m := map[string]interface{} {
    "w": &MyWriter{},
  }
}

func doSomething(m map[string]interface{}) {
  w := m["w"].(io.Writer) // This generates a runtime exception
}
```

This code compiles just fine. And if your test coverage is thorough, it might even pass that, too. But something is wrong.

```go
/**
Have the compiler do a type assertion for you
Error: cannot use MyWriter literal type (type *MyWriter)
*/
func TestWriter(t *testing.T) {
  var _ io.Writer = &MyWriter{}
}
```

The test fails because your `Write` method doesn't match the signature of `io.Writer`'s `Write([]byte) (int, error)`. The compilation error tells you exactly how to fix your writer to match the interface you intended to match.
