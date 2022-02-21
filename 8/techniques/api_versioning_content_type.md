# API version in content type

## Problem

How can API versions be handled in a semantic REST manner?

## Solution

Instead of referencing JSON in the request and response, use a custom content type that includes the version.

For example, instead of working with `application/json`, use a custom content type such as `application/vnd.mytodo.v1.json` or `application/vnd.mytodo.json; version=1.0`. These custom types specify the inteded schema for the data.

> See [API example](../api_versioning_content_type.go) and [requestor example](../request_api_versioned_with_ct.go).

## Discussion

To handle multiple API versions at a single path, the handling needs to take into account the content type in addition to any other characteristics.

Part of the riginal theory of REST was that a single URL represents a resource or group of resources, we don't want to have an api-specific path with the API version in the URL.
