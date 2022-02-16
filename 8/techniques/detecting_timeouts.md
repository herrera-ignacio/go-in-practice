# Detecting timeouts

- [Detecting timeouts](#detecting-timeouts)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

How can network timeouts be reliably detected?

## Solution

When timeouts occur, a small variety of errors occurs. Check the error for each of these cases to see if it was a timeout.

> [See example](../detecting_timeout.go)

## Discussion

When an error is returned from a `net` package operation or a package that takes advantage of `net`, such as `http`, check the `error` againast known cases showing a timeout error. Some of these will be for the explicit cases where a timeout was set and cleanly detected. Others will be for the cases where a timeout wasn't set but a timeout occurred.
