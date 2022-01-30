# Error Variables

- [Error Variables](#error-variables)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

One complex function may encounter more than one kind of error and we want to indicate which was returned to handle each error case appropraitely.

Although distinct error conditions may occur, none of them needs extra information.

## Discussion

## Solution

One convention is to *create package-scoped error variables* that can be returned whenever a certain error occurs.

> For example, the `io` package contains errors such as `io.EOF` and `io.ErrNoProgress`.
