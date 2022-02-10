# Extending templates with functions

- [Extending templates with functions](#extending-templates-with-functions)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

The built-in functions in the templates don't provide all the functionality you need.

## Solution

Just as Go makes funcitons available in templates (e.g., `fmt.Sprintf` as `printf`), make your own functions available.

> See [example](../template_functions/)

## Discussion

You can display information in templates in various ways. Although the way data is generated should be kept in the application logic, the way it's formatted and displayed should cleanly happen in a template.
