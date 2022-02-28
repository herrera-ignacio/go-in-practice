# Discovering whether a value implements an interface

- [Discovering whether a value implements an interface](#discovering-whether-a-value-implements-an-interface)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

Given a particular type, you want to find out whether that type implements a defined interface.

## Solution

There are two ways to accomplish this. One is with a type assertion, and the other uses the `reflect` package. Use the one that berst meets your needs.

> See [checking and converting a type](../checking_and_converting_type.go) and [check for interface](../check_for_interface.go) examples.

## Discussion

Go's view of interfaces differs from that of object-oriented languages like Java. In Go, a thing isn't declared to fulfill an interface. Instead, an interface is a description against which a type can be compared. And interfaces are themselves types.
