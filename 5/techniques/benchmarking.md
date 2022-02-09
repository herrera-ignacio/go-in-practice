# Benchmarking Go code

## Problem

You have code paths for accomplishing something, and you want to know which way is faster.

> It is faster to use `text/template` for formatting text, or just stick with `fmt`?

## Solution

Use the benchmarking feature, `testing.B`, to compare the two.

> See the [example](../benchmarking).

## Discussion

Benchmarks are treated similarly to tests. They go in the same `_test` files that unit tests and examples go in, and they're executed with the `go test` command.
