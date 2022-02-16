# Verify uploaded file is allowed type

- [Verify uploaded file is allowed type](#verify-uploaded-file-is-allowed-type)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How can you detect the type of file uploaded to a file field inside your application?

## Solution

To get the MIME type for a file, you can use one of a few ways with varying degrees of trust:

* When a file is uplodaded, use the request headers `Content-Type`.
* Use the file extension which is associated with a MIME type.
* Parse the file and detect the content type based on its contents.
