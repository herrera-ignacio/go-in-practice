# Reusing connections

- [Reusing connections](#reusing-connections)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)
    - [Problems preventing connection reuse](#problems-preventing-connection-reuse)

## Problem

When each request is over its own connection, a significant amount of time is lost to network communication. How can an application avoid as much of this lost time as possible?

## Solution

Reuse connections. Multiple HTTP requests can be made over a single connection, which needs to be negotiated and ramped up for slow-start only once.

> After passing the first message ,others happen more quickly.

## Discussion

It's not unusual for each HTTP request tobe made over its own connection. Negotiating each connection takes time, including the time to negotiate TLS for secure communications. Next, TCP slow-start ramps up as the message is communicated.

> _Slow-start_ is a congestion-control strategy designed to prevent network congestion. As a slow-start tramps up, a single message may take multiple round-trips between the client and the server to communicate.

Go tries to reuse connections out of the box, and it's the patterns in an application's code that can cause this to not happen. The server included in the `net/http` package provides HTTP keep-alive support.

### Problems preventing connection reuse

1. Custom transport instances are used and keep-alive is turned off.
2. Body of a response isn't close. Prior to HTTP/2, one request and response needed to be completed before the next cloud be used.
