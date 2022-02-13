# Mapping data types to templates

- [Mapping data types to templates](#mapping-data-types-to-templates)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

You want to render an object to HTML and pass the rendered object to a higher-level template, where it can be part of the output.

## Solution

Use templates to render objects as HTML. Store the HTML in a variable and pass the HTML to higher-level templates wrapped in `template.HTML`, marking it as safe HTML that doesn't need to be scapped.

> See [example](../mapping_data_types).
