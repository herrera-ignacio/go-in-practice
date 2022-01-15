# Locking wih a mutex

- [Locking wih a mutex](#locking-wih-a-mutex)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

Multiple goroutines are working with the same piece of data, and that data may change, you have the **potential for a race condition**.

## Solution

One simple way to avoid this situation is for each goroutine to place a _"lock"_ on a resource that it's using, and then unlock the resource when it's done. For all other goroutines, they wait until the resource is unlocked before attempting to lock that resource on their own. Use `sync.Mutex` to lock and unlock the resource.

## Discussion

The built-in sync package provides a `sync.Locker` interface as well as a couple of lock implementations. These provide essential locking behavior.

> Go includes built-in race detection in many of its Go tools by using a `--race` flag.
