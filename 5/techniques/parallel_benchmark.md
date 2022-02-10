# Parallel Benchmarks

- [Parallel Benchmarks](#parallel-benchmarks)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You want to test how a given piece of code performs when spread over goroutines. Ideally, you want to test this with a variable number of CPUs.

## Solution

A `*testing.B` instance provides a `RunParallel` method for exactly this purpose. Combined with command-line flags, you can test how well goroutines parallelize.

## Discussion

You can fashion the tempalte test into a parallel test. Instead of executing the body of a loop, the testing framework executes a function repeatedly, but as separate goroutines.

See [example](../parallel_benchmarks).
