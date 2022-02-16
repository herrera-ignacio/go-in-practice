# Serving from an alternative location

- [Serving from an alternative location](#serving-from-an-alternative-location)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

Instead of serving files through the same server as your application, you want to serve files through an alternative location. The alternative location needs to work with multiple environments such as your production, testing, and development environments.

## Solution

Serve the files from alternative locations, such as a CDN in prdoduction. For each environment, manage the deployment of files alongside the application and pass the location into the application as configuration. Use the location within your template files to tell browsers where to get the files from.

> See [example](../passing_url_to_template.go)
