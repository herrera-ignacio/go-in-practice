# Recovering from panics

- [Recovering from panics](#recovering-from-panics)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

A funciton your application calls is panicking, and as a result your program is crashing.

## Solution

Use a **deferred function** and call `recover` to find out what happened and handle the panic.

> See [Recovering from panics example](../recovering_from_panics.go) for more details.

## Discussion

Go provides the `recover` function as a way of capturing information from a panic and, stopping the panic from unwinding the function stack further.
