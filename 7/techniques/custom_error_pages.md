# File server with custom error pages

- [File server with custom error pages](#file-server-with-custom-error-pages)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How can you specify your own error pages, including a response to a file not being found, when your application is serving files?

## Solution

Use a custom file server that allows you to specify handlers for error pages. The `github.com/Masterminds/go-fileserver` package provides functionality to complement the built-in file server enabling custom error handling.

> See [example](../file_not_found.go).
