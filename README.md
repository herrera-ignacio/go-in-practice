# Go In Practice

- [Go In Practice](#go-in-practice)
  - [1. Noteworthy aspects of Go](#1-noteworthy-aspects-of-go)
  - [2. A solid foundation](#2-a-solid-foundation)
  - [3. Concurrency](#3-concurrency)
    - [Objectives](#objectives)
    - [Techniques](#techniques)
    - [Examples](#examples)
  - [4. Handling errors and panics](#4-handling-errors-and-panics)
    - [Objectives](#objectives-1)
    - [Techniques](#techniques-1)
    - [Examples](#examples-1)
  - [5. Debugging and testing](#5-debugging-and-testing)
    - [Objectives](#objectives-2)
    - [Techniques](#techniques-2)
    - [Examples](#examples-2)
  - [6. HTML and email template patterns](#6-html-and-email-template-patterns)
    - [Objectives](#objectives-3)
    - [Techniques](#techniques-3)
    - [Example](#example)
  - [7. Serving and receiving assets and forms](#7-serving-and-receiving-assets-and-forms)
    - [Objectives](#objectives-4)
    - [Techniques](#techniques-4)
    - [Example](#example-1)
  - [Working with web services](#working-with-web-services)
    - [Objectives](#objectives-5)
    - [Techniques](#techniques-5)
    - [Examples](#examples-3)
  - [9. Using the cloud](#9-using-the-cloud)
    - [Objectives](#objectives-6)
    - [Techniques](#techniques-6)
    - [Examples](#examples-4)
  - [10. Communication between cloud services](#10-communication-between-cloud-services)
    - [Objectives](#objectives-7)
    - [Techniques](#techniques-7)
    - [Examples](#examples-5)
  - [11. Reflection and code generation](#11-reflection-and-code-generation)
    - [Objectives](#objectives-8)
    - [Techniques](#techniques-8)
    - [Examples](#examples-6)

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

* Understanding Go's CSP-based concurrency model
* Concurrent processing with goroutines
* Locking and waiting with the `sync` package and *buffered channels*.
* Communication between goroutines using channels
* Strategically closing channels

### Techniques

1. [Using goroutines closures](3/techniques/closures.md)
2. [Waiting for goroutines](3/techniques/waiting.md)
3. [Locking with a mutex](3/techniques/mutex.md)
4. [Using multiple channels](3/techniques/multiple_channels.md)
5. [Closing channels](3/techniques/closing-channels.md)
6. [Locking with buffered channels](3/techniques/buffered_channels.md)

### Examples

1. [Using a goroutine to run a task](3/goroutine.go)
2. [Anonymous functions & closures](3/anonymous_function.go)
3. [Gzip compression tool](3/gzip_compression.go)
4. [Gzip compression with wait group](3/gzip_wait_group.go)
5. [Word counter](3/word_counter)
   1. [Word counter w/race condition](3/word_counter/word_counter_race_condition.go)
   2. [Word counter w/locks](3/word_counter/word_counter_locks.go)
6. [Using multiple channels](3/multiple_channels.go)
7. [Pausing with `Sleep` and `After`](3/pausing.go)
8. [Improper channel close](3/improper_channel_close.go)
9. [Close from sender](3/close_from_sender.go)
10. [Close using a close channel](3/close_channel.go)

## 4. Handling errors and panics

### Objectives

* Understand Go's patterns for error handling
* Using error variables and custom error types
* Providing meaningful data with errors
* Handling panics
* Transforming panics into errors
* Error handling on goroutines

### Techniques

> An *error* indicates that a particular task couldn't be completed successfully. A *panic* indicates that a severe event ocurred, probably as a result of a programmer error.

1. [Minimize the nils](4/techniques/minimize-nils.md)
2. [Custom error types](4/techniques/custom-error-types.md)
3. [Error variables](4/techniques/error_variables.md)
4. [Issuing panics](4/techniques/panic.md)
5. [Recovering from panics](4/techniques/recovering_from_panics.md)
6. [Trapping panics on goroutines](4/tecnniques/trapping_panics_on_goroutines.md)

### Examples

1. [Returning an error](4/returning_error.go)
2. [Relying on good error handling](4/relying_on_good_error_handling.go)
3. [Parse error](4/parse_error.go)
4. [Error variables](4/error_variables.go)
5. [Error and panic](4/error_and_panic.go)
6. [Recovering from panics](4/recovering_from_panics.go)
7. [Handle panics on a goroutine](4/handle_panics_on_goroutine.go)

## 5. Debugging and testing

### Objectives

* Capturing debugging information
* Using a logger (e.g., network)
* Capturing stack traces
* Writing unit tests.
* Benchmarking and acceptance tests
* Performing basic generative testing
* Detecting race conditions

### Techniques

1. [Logging to an arbitrary writer](5/techniques/logging_to_writer.md)
2. [Logging to a network resource](5/techniques/logging_to_network.md)
3. [Handling back pressure in network logging](5/techniques/back_pressure.md)
4. [Logging to the syslog](5/techniques/logging_to_syslog.md)
5. [Capturing stack traces](5/techniques/stack_traces.md)
6. [Using interfaces for mocking or stubbing](5/techniques/interfaces_for_testing.md)
7. [Verifying interfaces with canary tests](5/techniques/canary_tests.md)
8. [Generative testing](5/techniques/generative_testing.md)
9. [Benchmarking Go code](5/techniques/benchmarking.md)
10. [Parallel benchmarks](5/techniques/parallel_benchmark.md)
11. [Detecting race conditions](5/techniques/detecting_race_conditions.md)

### Examples

1. [Logging to a file](5/logging_to_file.go)
2. [Network log client](5/network_log_client.go)
3. [UDP-based logging](5/udp_logging.go)
4. [Logging to syslog](5/logger_to_syslog.go)
5. [Logging to system log](5/logger_to_system_log.go)
6. [Capturing stack traces with `runtime/debug`](5/capture_stack_trace.go)
6. [Capturing stack traces with `Stack` function](5/capture_stack_trace.go)
7. [Hello test](5/hello_test)
8. [Generative test](5/generative_test/)
9. [Benchmarking](5/benchmarking)
10. [Two templates benchmark](5/two_templates_benchmark)
11. [Parallel benchmarks](5/parallel_benchmarks)

## 6. HTML and email template patterns

### Objectives

* Extending the functionality within templates with custom functions and pipping commands.
* Caching & buffering templates
* Template inheritance
* Nesting templates
* Mapping templates to objects, such as a user template for a user object, and rolling the templates up into a page-level output.
* Generating email output with templates

### Techniques

1. [Extending templates with functions](6/techniques/template_functions.md)
2. [Caching parsed templates](6/techniques/caching_parsed.md)
3. [Handling template execution failures](6/techniques/handling_template_failures.md)
4. [Nested templates](6/techniques/nested_templates.md)
5. [Template Inheritance](6/techniques/template_inheritance.md)
6. [Mapping data types to templates](6/techniques/mapping_data_types.md)
7. [Generating email from templates](6/techniques/generating_email_from_templates.md)

### Example

1. [Simple HTML template](6/simple_template)
2. [Template functions](6/template_functions)
3. [Cache parsed templates](6/cache_template)
4. [Buffer templates for error handling](6/buffered_template)
5. [Template inheritance](6/template_inheritance)
6. [Mapping data types to templates](6/mapping_data_types)
7. [Email template](6/email_template)

## 7. Serving and receiving assets and forms

### Objectives

* Uploading files to users from a Go server.
* Go helper functions for quick and easy access to form submissions.
* Form parsing.
* Multipart form handling.

### Techniques

1. [Serving subdirectories](7/techniques/subdirectories.md)
2. [File server with custom error pages](7/techniques/custom_error_pages.md)
3. [Caching file server](7/techniques/caching_file_server.md)
4. [Embedding files in a binary](7/techniques/embedding_files.md)
5. [Serving from an alternative location](7/techniques/serving_from_alternative_location.md)
6. [Accessing multiple values for a form field](7/techniques/multiple_values_for_form_field.md)
7. [Uploading a single file](7/techniques/uploading_single_file.md)
8. [Uploading multiple files](7/techniques/uploading_multiple_files.md)
9. [Verify uploaded file is allowed type](7/techniques/verify_type.md)
10. [Incrementally saving a file (mutlipart data)](7/techniques/incrementally_saving_file.md)

### Example

1. [http package file serving](7/http_package_file_serving.go)
2. [Serve file with custom handler](7/servefile.go)
3. [Serve subdirectories with `http.Dir` handler](7/serving_subdirectory/httpdir.go)
3. [Serve subdirectories with `path` handler](7/serving_subdirectory/path.go)
4. [File not found error handler](7/file_not_found.go)
5. [Cache serving](7/caching_file_server.go)
6. [Embedding files in a binary](7/embedded_files.go)
7. [Serving from alternative location](7/http_package_file_serving.go)
8. [Parsing a simple form from response](7/parsing_form_from_response.go)
9. [Parsing multiple values from form](7/parsing_multiple_values.go)
10. [A form with a single-value file-upload field](7/single_upload)
11. [Multiple file uploads](7/multiple_uploads)

## Working with web services

### Objectives

* Detecting network timeouts.
* Resuming downloads when timeouts occur.
* Parsing errors between API endpoints and client requestors.
* Parsing JSON, even when you don't know the schema ahead of time.
* Versioning REST APIs with url and content-type.

### Techniques

1. [Detecting timeouts](8/techniques/detecting_timeouts.md)
2. [Timing out and resuming with HTTP](8/techniques/timing_out_and_resuming.md)
3. [Custom HTTP error passing](8/techniques/custom_http_error_parsing.md)
4. [Reading custom errors](8/techniques/reading_custom_errors.md)
5. [Parsing JSON without knowing the schema](8/techniques/parsing_json_no_schema.md)
6. [API version in the URL](8/techniques/api_versioning.md)
7. [API version in content type](8/techniques/api_versioning_content_type.md)

### Examples

1. [A simple HTTP get](8/simple_get.go)
2. [Delete request with default http client](8/simple_delete)
3. [Simple custom HTTP client](8/custom_http_client.go)
4. [Detecting timeouts and resuming with HTTP](8/download_with_retries.go)
5. [Passing an error over HTTP](8/passing_errors_over_http.go)
6. [Custom JSON error response](8/custom_json_error_res.go)
7. [Convert HTTP response to an error](8/convert_http_response_to_err.go)
8. [Parsing JSON](8/parsing_json.go)
9. [Parsing JSON with no schema](8/parsing_json_no_schema.go)
10. [API versioning in the URL](8/api_versioning_url.go)
11. [API versioning in the content type](8/api_versioning_content_type.go)
12. [Request API with content type versioning](8/request_api_versioned_with_ct.go)

## 9. Using the cloud

### Objectives

* Working with multiple cloud providers avoiding vendor lock-in.
* Gathering information about the host.
* Compiling to various operating systems and architectures.
* Monitoring the Go runtime to detect issues and details about a running application.

### Techniques

1. [Working with multiple cloud providers](9/techniques/multiple_cloud_providers.md)
2. [Cleanly handling cloud provider errors](9/techniques/cloud_provider_errors.md)
3. [Gathering information on the host](9/techniques/gathering_information_on_the_host.md)
4. [Detecting dependencies](9/techniques/detecting_dependencies.md)
5. [Cross-compiling](9/techniques/cross_compiling.go)
6. [Monitoring the Go runtime](9/techniques/monitoring_go_runtime.md)

### Examples

1. [Simple file storage](9/simple_file_storage.go)
2. [Handling cloud provider errors](9/handling_errors.go)
3. [Gathering information on the host](9/host_information.go)
4. [Detecting dependencies](9/detecting_dependencies.go)
5. [Monitoring the Go runtime](9/monitoring_go_runtime.go)

## 10. Communication between cloud services

### Objectives

* Communications in a microservice architecture.
* Reusing connections to improve performance by avoiding repeated TCP slow-start, congestion-control ramp-ups and connection negotiations.
* Faster JSON mashaling and unmarshaling that avoids extra time spent reflecting.
* Communicating over RPC using gRPC.

### Techniques

1. [Reusing connections](10/techniques/reusing_connections.md)
2. [Faster JSON marshal and unmarshal](10/techniques/faster_json.md)
3. [Using protocol buffers](10/techniques/protocol_buffers.md)
4. [Communicating over RPC with protocol buffers](10/techniques/rpc.md)

### Examples

1. [Faster JSON marshal and unmarshal](10/faster_json.go)
2. [Protocol buffer file](10/protocol_buffer_user.go)
3. [Protocol buffer server](10/user.proto)
4. [Protocol buffer client](10/protocol_buffer_client.go)
5. [RPC proto file](10/hello.proto)
6. [gRPC server](10/grpc_server.go)

## 11. Reflection and code generation

### Objectives

* Use kinds to identify critical details about types.
* Determine at runtime whether a type implements an interface.
* Access struct fields at runtime.
* Work with annotations.
* Parse tags within struct annotations.
* Write marshal and unmarshal functions.
* Use `go generate`.
* Write Go templates that generate Go code.

### Techniques

1. [Switching based on type and kind](11/techniques/switching_based_on_type_and_kind.md)
2. [Discovering whether a value implements an interface](11/techniques/checking_for_interface.md)
3. [Accessing fields on a struct](11/techniques/accessing_fields_on_a_struct.md)
4. [Processing tags on a struct](11/techniques/processing_tags.md)
5. [Generating code with go generate](11/techniques/go-generate.md)

### Examples

1. [Sum with type switch](11/sum_with_type_switch.go)
2. [Sum with Kind switch](11/sum_with_kind_switch.go)
3. [Checking and converting a type](11/checking_and_converting_type.go)
4. [Checking for an interface](11/check_for_interface.go)
5. [Recursively examining a value](11/examining_value.go)
6. [Simple JSON struct](11/json_struct.go)
7. [Processing tags](11/processing_tags.go)
8. [The queue template](11/queue_template.go) & [desired output](11/simple_queue.go)
