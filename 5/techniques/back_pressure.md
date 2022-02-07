# Handling back pressure in network logging

- [Handling back pressure in network logging](#handling-back-pressure-in-network-logging)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

Network log services are prone to connection failures and back pressure. This leads to lost log messages and sometimes even service failures.

## Discussion

You're likely to run into two major networking issues:

* The logger's network connectio ndrops.
* The connection slows down.

With TCP, the log server must send an ACK for each new message that comes in. When it delays, the client must slown down too, as its resources are tied up waiting for log messages to send. This is known as **back pressure**.

## Solution

Build a more resilient logger that buffers data. One solution to the back pressure problem is to switch from TCP to UDP.

> [See example](../udp_logging.go)

UDP for logging has distinct advantages:

* The app is resistant to back pressure and log server outages. If the log server hiccups, it may lose some UDP packets, but the client won't be impacted.
* Sending logs is faster even without back pressure.
* The code is simple.

But this route has some major disadvantages:

* Log messages can get lost easily because there's no message acknowledgement.
* Log messages can be received out of order. Adding a timestamp can help with this.
* Sending the remote server lots of UDP messages may turn out to be more likely to overwhelm the remote server.
