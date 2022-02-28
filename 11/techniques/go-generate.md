# Generating code with go generate

- [Generating code with go generate](#generating-code-with-go-generate)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You want to be able to create type-specific collections, such as a queue, for an arbitrary number of types. And you'd like to do it without the runtime safety issues and performance hit associated with type assertions.

## Solution

Build a generator that can create queues for you, and then use generation headers to generate the queues as you need them.

> See [desired output](../simple_queue.go) [queue template](../queue_template.go) examples.

## Discussion

Writing custom type-safe collections, gerating structs from database tables, transforming JSON schemata into code, generating many similar objects, and so on, is a common usage of Go generators.

Sometimes Go developers use the _Abstract Syntax Tree (AST)_ package or `yacc` tool to generate Go code. But sometimes, it's a good enough and easy way to build code by using Go templates that generate Go code.
