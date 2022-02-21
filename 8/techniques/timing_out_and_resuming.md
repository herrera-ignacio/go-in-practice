# Timing out and resuming with HTTP

- [Timing out and resuming with HTTP](#timing-out-and-resuming-with-http)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You want to resume downloading a file, starting from the end of the data already downloaded, after a timeout occurs.

## Solution

Retry the download again, attempting to use the `Range` HTTP header in which a range of bytes to download is specified. This allos wyou to request a file, starting partway through the file where it left off.

> See [example](../downloading_with_retries.go).

## Discussion

Servers, such as the one provided in the Go standard library, can support serving parts of a file. This is a fairly common feature in file servers.
