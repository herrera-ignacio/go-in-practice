# Using interfaces for mocking or stubbing

- [Using interfaces for mocking or stubbing](#using-interfaces-for-mocking-or-stubbing)
  - [Problem](#problem)
  - [Solution](#solution)
  - [Discussion](#discussion)

## Problem

You're writing code that depends on types defined in external libraries, and you want to write test code that can verify that those libraries are correctly used.

## Solution

Create interfaces to describe the types you need to test. Use those interfaces in your code, and then write stub or mock implementations for your tests.

## Discussion

Say you're writing software that uses a third-party library that describes some kind of message-sending system:

```go
type Message struct {
  // ...
}

func (m *Message) Send(email, subject string, body []byte) error {
  // ...
  return nil
}
```

In your code, you use that library to send a message from your application. You want to ensure that the code that sends the message is being called, but you don't want to send the message.

One way to deal with this is to write your own interface and have your code use it in its declarations instead of directly using the `Message` type.

```go
type Messager interface {
  Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
  return m.Send("test@example.com", "Critical error", problem)
}
```

Because you've created an abstraction from `Message` to `Messager`, you can easily write a mock and use that for your testing.

```go
package msg

import "testing"

type MockMessage struct {
  email, subject string
  body           []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error {
  m.email = email
  m.subject = subject
  m.body = body
  return nil
}

func TestAlert(t *testing.T) {
  msgr := new(MockMessage)
  body := []byte("Critical Error")

  Alert(msgr, body)

  if msgr.subject != "Critical Error" {
    t.Errorf("Expected 'Critical Error', got '%s'", msgr.subject)
  }
}
```
