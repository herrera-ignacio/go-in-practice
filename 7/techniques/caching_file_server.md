# Caching file server

## Problem

Instead of serving static files from the filesystem each time they're requested, youw ant to cache files in memory in order to quickly serve resposnes to requests.

## Solution

Store files in memory when they're first requested and serve responses using `Serve-Content` rather than a file server.

Most of the time, it's appropriate to use a reverse proxy, such as the popular open source project _Varnish_, to handle caching and serving files of quickly.

> See [example](../caching_file_server.go)
