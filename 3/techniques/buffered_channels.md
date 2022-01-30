# Locking with buffered channels

- [Locking with buffered channels](#locking-with-buffered-channels)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You need to lock certain resources. Given the frequent use of channels in your code, you'd like to do this with channels instead of the `sync` package.

If an *unbuffered channel* - a channel that that contain one value at a time and are created like `make(chan TYPE)` - has received a value, and is then sent another one before the channel can be read, the second send operation will block. Moreover, the sender will also block until the channel is read.

## Solution

Use a channel with a buffer size of 1, and share the channel among the goroutines you want to synchronize. 

## Discussion

> The `sync` package is part of Go's core, and is well tested and maintained. You could use `sync.Locker` and `sync.Mutex`. But sometimes, especially in code that already uses channels, it's desirable to implement locks with channels instead of the mutex. Often this is a stylistic preference; it's prudent to keep your code as uniform as possible.

When talking about using a channel as a lock, you want this kind of behavior:

1. A function acquires a lock by sending a message on a channel.
2. The function proceeds to do its sensitive operations.
3. The function releases the lock by reading the message back off the channel.
4. Any function that tries to acquire the lock before it's been released will pause when it tries to acquire the (already locked) lock.

> In other words, the sender waits until something receives the message it puts on the channel.

One of the features of a *buffered channel* is that it doesn't block on send provided that buffer space still exists. A sender can send a message into a buffer and then move on. But if a buffer is full, the sender will block until there's room in the buffer for it to write its message.

If you create a channel with only one empty buffer space, one function can send a message, do its thing, and then read the message off the buffer (thus unlocking it). **This is exactly the behavior you want in a lock**.
