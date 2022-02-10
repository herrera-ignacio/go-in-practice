# Detecting race conditions

## Problem

In programs with many goroutines, race conditions could occur. Being able to test for this possibility is desirable.

## Solution

Use the `-race` flag (sometimes called `go race` or `grace`).

## Discussion

Let's make an ill-conceived performance optimization, see [example](../detecting_race_conditions).
