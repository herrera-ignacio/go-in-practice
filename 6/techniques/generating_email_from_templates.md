# Generating email from templates

- [Generating email from templates](#generating-email-from-templates)
  - [Problem](#problem)
  - [Solution](#solution)

## Problem

When creating and sending email, you want to incorporate templates.

## Solution

Use the template packages to generate the email text into a buffer. Pass the generated email in the buffer to the code used to send the email, such as the `smtp` package.

> See [example](../email_template)
