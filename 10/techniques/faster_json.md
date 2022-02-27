# Faster JSON marshal and unmarshal

- [Faster JSON marshal and unmarshal](#faster-json-marshal-and-unmarshal)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

Instead of figuring out the types of data each time JSON is marshaled or unmarshaled, how can the type be figured out once and skipped on future passes?

## Solution

Use a package able to generate code that can marshal and unmarshal de JSON. The gneerated code skips reflection and provides a faster execution path with a smaller memory footprint.

> See [example using ugorji/go/codec](../faster_json.go)

## Discussion

The JSON marshaling and unmarshaling provided by `encoding/json` package uses reflection to figure out values and types each time. Reflection, provided by the `reflect` package, takes time to figure out types and values each time a message is acted on. If you're repeatedly acting on the same structures, quite a bit of time will be spent reflecting.
