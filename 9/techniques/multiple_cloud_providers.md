# Working with multiple cloud providers

- [Working with multiple cloud providers](#working-with-multiple-cloud-providers)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

Cloud service providers typically have their own APIs, even when they offer the same or similar features. Writing an application to work with one API can lead to _lock-in_ with that API and that vendor.

## Solution

First, create an interface to describe your cloud interactions. If you need to save a file or create a server instance, have a definition on the interface for that.

Second, create an implementation for each cloud provider you're going to use.

## Discussion

The following listing provides an example of one designed to work with files.

```go
type File interface {
  Load(string) (io.ReadCloser, error)
  Save(string io.ReadSeeker) error
}
```

Then you need a first implementation. The easiest one, which allows you to test and locally develop your application, is to use the local filesystem as a store. This allows you to make sure the application is working before introducing network operations or a cloud provider.

> See [example](../simple_file_storage.go).
