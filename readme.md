# templatereq
A Go library to replace strings in a template using custom functions or map values. It supports different hash functions, base64 encoding, date manipulation, and more. This library uses regex to replace the values in a given string based on the provided map. It also provides a variety of string manipulation functions.

## Features
- Replace string based on map values
- Hash functions
- Base64 encoding
- SHA256
- MD5
- Date formatting and adjustments
- Lowercase and uppercase conversion with optional encryption
- UUID generation
- URL encoding

## Getting Started
### Prerequisites
- Go version 1.13 or higher

### Installing
```
go get github.com/<your_username>/templatereq
```

## Usage

```go
import (
    "github.com/<your_username>/templatereq"
)

// Create a map for the template replacements
var values = map[string]string{
    "TEST": "HELLO",
}

var template = "https://www.testweb.com/$TEST"
// Use Replace function to replace the template values
result := templatereq.Replace(template, values)
fmt.Println(result) // Output: https://www.testweb.com/HELLO
```

## Functions

Here is a list of the available functions:

- `$func("hash:<value>")`: Hashes a value using the fnv32a hash algorithm.
- `$func("md5:<value>")`: Returns the md5 hash of a value.
- `$func("base64:<value>")`: Returns the base64 encoding of a value.
- `$func("sha256:<value>")`: Returns the sha256 hash of a value.
- `$func("dateFormat:<value>")`: Formats a date value.
- `$func("lowercase:<value>")`: Converts a value to lowercase.
- `$func("uppercase:<value>")`: Converts a value to uppercase.
- `$func("uuid")`: Generates a new UUID.

## Running the tests

The `replace_test.go` file contains tests for this library. You can run them using the command `go test`.

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.