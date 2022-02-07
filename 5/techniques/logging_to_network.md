# Logging to a network resource

- [Logging to a network resource](#logging-to-a-network-resource)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

Streaming logs to a network service is error-prone, but you don't want to lose log messages if you can avoid it.

## Discussion

Logging to a network offers compelling advantages:

* Logs from many services can be aggregated to one central location.
* In the cloud, servers with only ephemeral storage can still have logs perserved.
* Security and auditability are improved.
* You can tune log servers and app servers differently.

But there's **one major drawback** with remote logging: **you're dependent on the **network**.

## Solution

Use Go's channels and some buffering to improve reliability.

> See [example](../network_log_client.go).
