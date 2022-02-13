# Template Inheritance

- [Template Inheritance](#template-inheritance)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

You want to have a base template and have other templates extend the base template. The templates would have multiple sections that can be extended.

## Solution

Instead of thinking of a file as a template, think of sections of a file as templates. The base file contains the shared markup and refers to o-ther templates that haven't yet been defined. The templates extending the base file provide the missing subtemplates or override those in the base.

The template system enables some inheritance patterns within templates.

> See [example](../template_inheritance)
