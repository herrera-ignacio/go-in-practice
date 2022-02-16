## Embedding files in a binary

- [Embedding files in a binary](#embedding-files-in-a-binary)
- [Problem](#problem)
- [Solution](#solution)
- [Discussion](#discussion)

## Problem

You want to include static assets, such as images or stylesheets, in a Go binary.

## Solution

Store the assets as data assigned to variables in your application. Serve files from these variables instead of the filesystem. Because of the intricate nature of converting files' bytes and referencing them within your code, use the `github.com/GeertJohan/go.rice` package and command-line application to automate this process for you.

> See [example](../embedded_files.go)

## Discussion

The idea is simple. Convert a file into bytes, store those bytes, and the related information in a variable, and use the variable to serve the file via `ServeContent` from the `http` package.

