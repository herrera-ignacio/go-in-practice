# Reading custom errors

- [Reading custom errors](#reading-custom-errors)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

When a custom error with a different structure is returned as an API response, how can you detect that and handle it differently?

## Solution

When a response is returned, check the HTTP status code and MIME type for a possible error. When one of these returns unexpected values or informs of an error, convert it to an `error`, return it, and handle it.

Go is known for **explicit error handling**.

> See [example](../convert_http_response_to_err.go).
