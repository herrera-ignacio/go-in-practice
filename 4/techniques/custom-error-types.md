# Custom error types

- [Custom error types](#custom-error-types)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

Errors returned by function might lead users of this function to code differently, depending on its details.

## Discussion

Imagine you're writing a file parser. When the parser encounters a syntax error, it generates an error. Along with having an error message, it'
s generally useful to have information about where in the file the error ocurred.

## Solution

Create a type that implements the error interface but provides additional functionality.

> See [example](../parse_error.go)
