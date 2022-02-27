# Using protocol buffers

- [Using protocol buffers](#using-protocol-buffers)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

What formats that are optimized for network transfer and serialization are available for use in Go applications?

## Solution

Some more recent formats, including _protocol buffers_ (i.e., _protobuf_), are faster in network transfer and serialization.

## Discussion

Protocol buffers are a popular choice for a high-speed transfer format. The data in the messages being transferred over the network is smaller than XML or JSON and it can be marshaled and unmarshaled faster as well.

The transfer method isn't defined by protocol buffers.

A protocol format is defined in a file. It contains the structure of the message and can be used to automatically generate the needed code.

> See [protocol buffer file](../user.proto), [protocol buffer server](../protocol_buffer_server.go) and, [protocol buffer client](../protocol_buffer_client.go) examples.

You'll need to compile the protocol buffer to code.

1. Download and install the compiler from https://developers.google.com/protocol-buffers
2. Install the Go protocol buffer plogin: `go get -u github.com/golang/protobuf/proto`.
3. Generate the code running `protoc -I=. --go_out=. ./user.proto`.
