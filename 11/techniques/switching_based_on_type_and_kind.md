# Switching based on type and kind

- [Switching based on type and kind](#switching-based-on-type-and-kind)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You want to write a function that takes generic values (`interface{}`s), and then does something useful with them based on underlying types.

## Solution

Go provides various methods for learning this information, ranging from the type switch to the `reflect.Type` and `reflect.Kind` types. Each has subtle strong points.

## Discussion

Say you want to write a function with the signature `sum(...interface{}) float64`. You want this function to take any number of arguments of various types. And you want it to convert the values to `float64` and then sum them.

> See [sum with type switch example](../sum_with_type_switch.go).

In a type switch, if you have custom types that match an underlying kind (e.g., `type MyInt int64`), it won't match when asserting for the underlying kind (i.e., `MyInt` won't be matched as `int64`).

The solution to this problem is to use the `reflect` package.

> See [sum with `Kind` switch](../sum_with_kind_switch.go)
