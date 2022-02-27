# Cleanly handling cloud provider errors

- [Cleanly handling cloud provider errors](#cleanly-handling-cloud-provider-errors)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How can application code handle interchangeable implementations of an interface while displaying or acting on errors returned by methods in the interface?

## Solution

Along with the methods on an interface, define and export errors. Instead of returning implementation-specific errors, return package errors.

> See [example](../handling_errors.go)
