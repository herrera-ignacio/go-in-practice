# Generative Testing

- [Generative Testing](#generative-testing)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You want to bulletproof your code against surprising edge cases.

> *Generative testing* is a large and complex topic. But in its most basic form, generative testing refers to the strategy of automatically generating test data in order to both broaden the information tested and overcome our biases when we choose our test data.


## Solution

Use Go's `testing/quick` package to generate testing data.

> See [example](../generative_test).

## Discussion

The `testing/quick` package provides several helpers for rapidly building tests that are more exhaustive than usual.
