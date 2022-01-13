# Using goroutine closures

- [Using goroutine closures](#using-goroutine-closures)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

> Go allows you to declare functions inline, you can **share variables** by declaring one function inside another and closing over the variables you want to share.

## Problem

You want to use a one-shot function in a way that doesn't block the calling function, and you'd like to make sure that it runs. For example, read a file in the background, send messages to a remote log server, or save a current state without pausing the program.

## Solution

Use a closure function and give the scheduler opportunity to run.

## Discussion

In Go, functions are _first-class_. They can be created inline, passed into other functions, and assigned as values to a variable. You can even declare an anonymous function and call it as a goroutine.
