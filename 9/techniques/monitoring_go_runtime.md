# Monitoring the Go runtime

- [Monitoring the Go runtime](#monitoring-the-go-runtime)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How can your application log or otherwise monitor the Go runtime?

## Solution

The `runtime` and `runtime/debug` packages provide access to the information within the runtime. Retrieve information from the runtime by using these packages and write it to the logs or other monitoring service at regular intervals.

> See [example](../monitoring_go_runtime.go)
