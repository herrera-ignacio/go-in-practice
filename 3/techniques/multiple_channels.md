# Using multiple channels

- [Using multiple channels](#using-multiple-channels)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You want to use channels to send data from one goroutine to another, and be able to interrupt that process to exit.

## Solution

Sometimes the best way to solve concurrency problems in Go is to communicate more information. And that ofen translates into using more channels.

Use `select` and multiple channels to signal when something is done or ready to close.

## Discussion

The `select` statement can watch multiple channels (zero or more). Until something happens, it'll wait (or execute a `default` statement, if supplied). When a channel has an event, the `select` statement will execute that event.
