# Detecting dependencies

- [Detecting dependencies](#detecting-dependencies)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

How can you ensure that it's okay to execute an application before calling it?

## Solution

Prior to calling a dependent application for the first time, detect whether the application is installed and available for you to use. Otherwise, log an error to help with troubleshooting.

> See [example](../detecting_dependencies.go)
