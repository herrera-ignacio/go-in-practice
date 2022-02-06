# Logingg to writer

- [Logingg to writer](#logingg-to-writer)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

You want to send logging messages to a file or to a network service without having to write your own logging system.

## Discussion

The `log.Logger` provides features for sending log data to any `io.Writer`, which includes things like file handles and network connections (`net.Conn`).

## Solution

Initialize a new `log.Logger` and send log messages to that.

> See [example](../logging_to_file.go).
