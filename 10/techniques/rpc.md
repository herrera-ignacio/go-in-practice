# Communicating over RPC with protocol buffers

- [Communicating over RPC with protocol buffers](#communicating-over-rpc-with-protocol-buffers)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)
    - [`context`](#context)
    - [Advantages](#advantages)
    - [Disadvantages](#disadvantages)

## Problem

How can you communicate over RPC in a manner that's protable across programming languages?

## Solution

Use gRPC and protocol buffers to handle defining the interfaces, generating the cross-language code, and implementing the RPC communications.

> See [RPC interface](../hello.proto) [gRPC server](../grpc_server.go) and [gRPC request](../grpc_request.go) examples.

You need to generate code for proto file: `protoc -I=. --go_out=plugins=grpc:. ./hello.proto`.

## Discussion

gRPC is an open source, high-performance RPC framework that can use HTTP/2 as a transport layer. It was developed at Google and has support for several programming languages.

gRPC uses code generation to create messages and handle parts of the communication. This enables the communication to easily work between multiple languages, as the interfaces and messages are generated properly for each language. 

### `context`

The `context` package can be used for several use cases. One example is using a context with a cancel function. This is useful if the caller goes away (for example, the application is closed), and the callee needs to be informed on this. The client could create a context like this:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

When the function this code is in ends, the `cancel()` function is called. This tells the context to cancel the work. This is communicated to the other service, even when hosted on an entirely different system. On the server, this cancellation can be seen through a channel, which is how it's implemented in the client as well.

```go
select {
  case <-ctx.Done():
    return nil, ctx.Err()
}
```

### Advantages

* Using protocol buffers, the payload size is smaller and faster to marshal/unmarshal than JSON or XML.
* The context allows for canceling, communicating timeouts or due dates, and other relevant information over the course of a remote call.
* The interaction is with a called procedure rather than a message that's passed elsewhere.
* Semanntics of the transport methodology (e.g., HTTP verbs) don't limit the communications over RPC.

### Disadvantages

* Transport payload isn't human readable as JSON or XML.
* Applications need to know the interface and details of the RPC calls. Knowing the semantics of the message isn't enough.
* Integration is deeper than a message, such as you'd have with REST, because a remote procedure is being called. Exposing remote procedure access to untrusted clients may not be ideal and should be handled with care for security.
