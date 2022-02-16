# Uploading a single file

- [Uploading a single file](#uploading-a-single-file)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

When a file is uploaded with a form, how do you process and save it?

## Solution

When a file is uploaded, process the form as a multipart form by using `Process-MultipartForm` on the `Request` object. This picks up the file parts. Then use the `FormFile` method on the `Request` object to access and file fields, uploading a single file. For each file, you can access the metadata and a file object that's similar to `File` objects from the `os` package.

> See [example](../uploading_single_file.go)

## Discussion

Handling a file is nearly as straightforward as handling text form data. The difference lies in the binary file and the metadata surrounding it, such as the filename.
