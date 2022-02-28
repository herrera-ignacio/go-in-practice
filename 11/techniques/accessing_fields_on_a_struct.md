# Accessing fields on a struct

- [Accessing fields on a struct](#accessing-fields-on-a-struct)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

You want to learn about a struct at runtime, discovering its fields.

## Solution

Reflect the struct and use a combination of `reflect.Value` and `reflect.Type` to find out information about the struct.

> See [recursively examining a value example](../examining_value.go)
