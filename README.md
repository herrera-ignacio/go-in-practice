# Go In Practice

- [Go In Practice](#go-in-practice)
  - [1. Noteworthy aspects of Go](#1-noteworthy-aspects-of-go)
  - [2. A solid foundation](#2-a-solid-foundation)
  - [3. Concurrency](#3-concurrency)
    - [Objectives](#objectives)
    - [Techniques](#techniques)
    - [Examples](#examples)

## 1. Noteworthy aspects of Go

1. [Multiple returns](1/returns.go)
2. [Named return values](1/returns2.go)
3. [Read TCP Status](1/read_status.go)
4. [HTTP Get](1/http_get.go)
5. [Concurrent output](1/concurrent_print.go)
6. [Using channels](1/channel.go)
7. [Hello world](1/hello.go)
8. [Testing hello world](1/hello_test.go)
9. [Hello World web server](1/nago.go)

## 2. A solid foundation

1. ["Hello" CLI w/`flag`](2/flag_cli.go)
2. ["Hello" CLI w/`go-flags`](2/go_flags/main.go)
3. ["Hello" CLI w/`cli.go`](2/cli_go/main.go)
4. [Count Up/Down w/`cli.go`](2/count_cli/main.go)
5. [Using JSON config](2/config_json/main.go)
6. [Using YAML config](2/config_yaml/main.go)
7. [Using INI config](2/config_ini/main.go)
8. [Using env variables](2/env_config.go)
9. [Callback shutdown URL (anti-pattern)](2/callback_shutdown.go)
10. [Graceful shutdown using manners](2/manners_shutdown/main.go)
11. [Resolve URLs with handler functions](2/multiple_handlers.go)
12. [Resolve URLs using `path`](2/path_handlers/main.go)
13. [Resolve URLs using RegExp](2/regex_handlers.go)
14. Faster routing
    1. [github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
    2. [github.com/gorilla/mux](https://github.com/gorilla/mux)
    3. [github.com/bmizerany/pat](https://github.com/bmizerany/pat)

## 3. Concurrency

### Objectives

* Understanding Go's concurrency model
* Using goroutines for concurrent processing
* Locking and waiting
* Using channels for communication between goroutines
* Strategically closing channels

### Techniques

1. [Using goroutines closures](3/techniques/closures.md)
2. [Waiting for goroutines](3/techniques/waiting.md)
3. [Locking with a mutex](3/techniques/mutex.md)

### Examples

1. [Using a goroutine to run a task](3/goroutine.go)
2. [Anonymous functions & closures](3/anonymous_function.go)
3. [Gzip compression tool](3/gzip_compression.go)
4. [Gzip compression with wait group](3/gzip_wait_group.go)
5. [Word counter](3/word_counter)
   1. [Word counter w/race condition](3/word_counter/word_counter_race_condition.go)
   2. [Word counter w/locks](3/word_counter/word_counter_locks.go)
