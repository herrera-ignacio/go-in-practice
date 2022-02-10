# Handling template execution failures

- [Handling template execution failures](#handling-template-execution-failures)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

When an error happens while executing a template, you want to catch the error before anything is displayed to the user. Instead of displaying the partially broken pages, display something more appropriate, such as an error page.

## Solution

Write the output of the template to a buffer. If no errors occur, write the output of the buffer to end users. Otherwise, handle the error.

> See [example](../buffered_template)

## Discussion

Any errors in the data should be handled before templates are used with the data, and the functions called within the templates should be for display purposes. This keeps the separation of concerns in place and limits failures when templates are executed, which is important for _streaming_.

Streaming responses is useful. When you execute a template to a response writer, end users start to receive the page more quickly. *When you buffer a response, there's a delay in end users receiving it*.

> End users expect native desktop performance out of web applications, and streaming responses help achieve that. When possible, write to output.

If executing templates carries a potential for errors, you can write the output of a template to a buffer so that if errors occur, they can be handled before displaying anything to the end users.
