# Trapping panics on goroutines

## Problem

If a panic on a goroutine goes unhandled, it crashes the entire program. Sometimes this is acceptable, but often the desired behavior is recovery.

## Solution

Simply stated, handle panics on any goroutine that might panic.

## Discussion

Although it's trivally easy to solve, the solution is repetitive, and often the burden of implementing it is pushed to developers outside your control.
