# Serving subdirectories

- [Serving subdirectories](#serving-subdirectories)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

You want to serve a directory and its subdirectories from the filesystem as part of your web application.

## Solution

Use the built-in file server of the file-serving handlers. For intimate control over error pages, including the case of a file not being found, you need to implement your own file server.

> See [`http.Dir` example](../serving_subdirectory/httpdir.go) and [`path` example](../serving_subdirectory/path.go).
