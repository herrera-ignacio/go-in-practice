# Uploading multiple files

- [Uploading multiple files](#uploading-multiple-files)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How do you process the files when multiple files are uploaded to a single file-input field on a form?

## Solution

Parse the form and retrieve a slice with the files from the `MultipartForm` property on the `Request`. Then iterate over the slice, individually handling each file. An input handling multiple files needs to have only the `multiple` attribute on it.

> See [example](../multiple_upload)
