# Processing tags on a struct

- [Processing tags on a struct](#processing-tags-on-a-struct)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

You want to create your own annotations and then programmatically access the annotation data data of a struct at runtime.

## Solution

Define your annotation format. Then use the `reflect` package to write a tool that extracts the annotation information from a struct.

> See [example](../processing_tags.go)
