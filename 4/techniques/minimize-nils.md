# Minimize the nills

- [Minimize the nills](#minimize-the-nills)
  - [Problem](#problem)
  - [Discussion](#discussion)
  - [Solution](#solution)

## Problem

Returning `nil` results along with errors isn't a great practice: it puts more work on your library's users, provides little useful information, and makes recovery harder.

## Discussion

Nils are annoying for several reasons:

1. Frequent cause of bugs because it forces to check values to protect against nils.
2. Sometimes they are used to indicate something specific (e.g., no errors, no data, no results, etc.).
3. Sommetimes they are treated as placeholders.

## Solution

When it makes sense, avail yourself of Go's multiple returns and send back not just an error, but also a usable value.

> When you're creating an error, you can use the `errors.New` function for creating simple new errors, and `fmt.Errof` for using a formatting string on the error message.
