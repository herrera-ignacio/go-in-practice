# Waiting for goroutines

- [Waiting for goroutines](#waiting-for-goroutines)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

> Sometimes you'll want to start multiple goroutines but not continue until they're all done. _Go wait groups_ are a simple way to achieve this.

## Problem

One goroutine needs to start one or more other goroutines, and then wait for them to finish. For example, you want to compress multiple files as fast as possible and then display a summary.

## Solution

Run individual tasks inside goroutines. Use `sync.WaitGroup` to signal the outer process that the goroutines are done and it can safely continue.

> Several workers are started, and work is delegated to the workers. One process delegates the tasks and then waits for the workers to complete.

## Discussion

Go's standard library provides several useful tools for working with synchronization. `sync.WaitGroup` comes in handy for telling one goroutine to wait until other goroutines complete.
