# Cross-compiling

- [Cross-compiling](#cross-compiling)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How can you compile for architectures and operating systems other than the one you're currently on?

## Solution

The `go` toolchain provides the ability to cross-compile. In addition, `gox` allows you to cross-compile multiple binaries in parallel. You also can use packages, such as `filepath` to handle differences between OS instead of hardcoding values, such as the POSIX path separator `/`.

```
gox \
  -os="linux darwin windows" \
  -arch="amd64 386" \
  -output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}"
```
