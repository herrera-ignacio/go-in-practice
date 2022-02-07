# Capturing stack traces

- [Capturing stack traces](#capturing-stack-traces)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

You want to fetch a stack trace at a critical point in the application.

## Solution

Use the `runtime package`, which has several tools.

If all you need is a trace for debugging, you can easily send one to *standard output* by using `runtime/debug` function `PrintStack`.

> See [example 1](../capture_stack_trace.go) and [example 2](../capture_stack_trace_2.go).
