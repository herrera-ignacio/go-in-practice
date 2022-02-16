# Accessing multiple values for a form field

- [Accessing multiple values for a form field](#accessing-multiple-values-for-a-form-field)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

`FormValue` and `PostFormValue` each return the first value for a form field. When you have multiple values, how can you access all of them?

## Solution

Lookup the fields on the `Form` or `PostForm` properties on the `Request` object. Then iterate over all the values.

> See [example](../parsing_multiple_values.go)
