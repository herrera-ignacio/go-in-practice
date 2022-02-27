# Gathering information on the host

- [Gathering information on the host](#gathering-information-on-the-host)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

How can information about a host be detected within a Go application?

## Solution

The `os` package enables you to get information about the underlying system, moreover, it can be combined with information detected through other packages, such as `net` or from calls to external applications.

> See [example](../host_information.go)

## Discussion

The `os` package has the capability to detect a wide range of details about the environment.

* `os.Hostname()` returns the kernel's value for the hostname.
* The process ID for the application can be retrieved wit `os.Getpid()`
* Operating systems can and do have different path and path list separators. Using the `os.PathSeparator` or `os.PathListSeparator` instead of characters allows applications to work with the system they're running on.
* Toi find the current working directory, use `os.Getwd()`.
