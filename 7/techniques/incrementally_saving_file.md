# Incrementally saving a file

- [Incrementally saving a file](#incrementally-saving-a-file)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

You want to save the file, as it's being uploaded, to a location of your choice. That location could be on the server, on a shared drive, or on another location altogether.

## Solution

Read the multipart data from the request as it's being uploaded. This can be accessed with the `MultipartReader` method on the `Request`. As files and other informations are coming in, chunk by chunk, save and process the parts rather than wait for uploads to complete.

> See [example](../multipart_upload)
