# Custom HTTP error passing

- [Custom HTTP error passing](#custom-http-error-passing)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How can you provide a custom response body and content type when there's an error?

## Solution

Instead of using the built-in `Error` function, use custom functions that send both the correct HTTP status code and the error text as a more appropriate body for your situation.

> See [example](../custom_json_error_res.go).
