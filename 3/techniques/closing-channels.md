# Clossing channels

## Problem

What happens if you have a sender and receiver goroutine, and the sender finishes sending data? Are the receiver goroutine and channel automatically cleaned up? Nope. The memory manager will only clean up values that it can ensure won't be used again. You have to be careful not to leak both channels and goroutines.

You don't want leftover channels and goroutines to consume resources and cause leaky applications. You want to safely close channels and exit gorotutines.

> Closing channels the wrong way will cause your program to panic or leak goroutines.

## Solution

The predominant method is to use additional channels to notify goroutines when it's safe to close a channel.

> "Close your channels and return from your goroutines."

As a rule of thumb, a receiver should never close a receiving channel by itself. Instead, it sends a message on a `done` channel indicating that it's done with its work. 

## Discussion

YOu can use a few idiomatic techniques for safely shutting down channels. See [3.8](../improper_channel_close.go), [3.9](../close_from_sender.go), and [3.10](../close_channel.go) examples.
