# Issuing panics

- [Issuing panics](#issuing-panics)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

When you raise a panic, what should you pass into the function? Are there ways of panicking that are useful or idiomatic?

## Discussion

What sort of thing would cause a panic? An error. It's reasonable to assume that developers will expect this. Also, fulfilling the error interface eases handling of a panic.

## Solution

**The best thing to pass to a panic is an error**. Use the error type to make it easy for the recovery function (if there is one).

```go
func main() {
  panic(errors.New("Something bad happened."))
}
```
