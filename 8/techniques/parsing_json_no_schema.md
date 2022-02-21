# parsing JSON without knowing the schema

## Problem

How can you parse a JSON data structure into a Go data structure when you don't know the structure ahead of time?

## Solution

Parse the JSON into an `interface{}` instead of a struct. Afterwards, you can inspect the data and use it.

> See [example](../parsing_json_no_schema.go)

## Discussion

A little-known feature of the `encoding/json` package is the capability to parse arbitrary JSON into an `interface{}`. It is quite different from working with JSON parsed into a known structure, because of the Go type system.
