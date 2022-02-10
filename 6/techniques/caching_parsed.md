# Caching parsed templates

- [Caching parsed templates](#caching-parsed-templates)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You want to avoid repeatedly parsing the same template while an application is running.

## Solution

Parse the template, store the ready-to-use template in a variable, and repeatedly use the same template each time you need to generate the output.

> See [example](../cache_template)

## Discussion

Instead of parsing the template each time in the `http` handler function, move the parsing out of the handler. Then you can repeatedly execute the template against different datasets without parsing it each time.

