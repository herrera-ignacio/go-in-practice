# Logging to the syslog

- [Logging to the syslog](#logging-to-the-syslog)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

You want to send applicatioon log messages into the system logger.

## Discussion

The `syslog` package gives you two ways of working with the system log:

1. Use it as a logging back end.
2. Use all of the defined log levels and facilities directly in a syslog-specific style.

## Solution

Configure Go's `syslog` package and use it.

> See [example 1](../logger_to_syslog.go) and [example 2](../logger_to_system_log.go).
