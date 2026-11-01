# initials

[![GoDoc](https://godoc.org/github.com/go-mods/initials?status.svg)](https://godoc.org/github.com/go-mods/initials)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mods/initials)](https://goreportcard.com/report/github.com/go-mods/initials)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Initials is a Go package that generates initials from names, for example to display an avatar.
Words are separated by spaces and hyphens, special characters are removed, and accented or
non-Latin letters listed in its character table are transliterated. Other characters, such as
emoji or CJK ideographs, are kept as they are.

<details>
<summary>Table of contents</summary>

- [Installation](#installation)
- [Usage](#usage)
- [Options](#options)
- [Development](#development)
- [License](#license)

</details>

## Installation

Prerequisites: Go 1.22 or later.

```shell
go get github.com/go-mods/initials
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/go-mods/initials"
)

func main() {
    fmt.Println(initials.GetInitials("John Doe"))                                                   // JD
    fmt.Println(initials.GetInitials("John Doe", initials.WithLength(3), initials.WithCamelCase())) // JDo
    fmt.Println(initials.GetInitials("John"))                                                       // JO
    fmt.Println(initials.GetInitials("John", initials.WithCamelCase()))                             // Jo
}
```

## Options

The initials are in uppercase by default. `WithSensitive`, `WithLowercase`, `WithUppercase` and
`WithCamelCase` are mutually exclusive: the last one given applies.

### WithLength (_default: 2_)

Sets the maximum number of characters. A length of zero or less gives an empty string.

```go
initials.GetInitials("John Doe", initials.WithLength(3)) // JDO
```

### WithWordLength

Sets the length to the number of words.

```go
initials.GetInitials("John David Dugga", initials.WithWordLength()) // JDD
```

### WithSeparator (_default: ""_)

Sets the separator between the initials.

```go
initials.GetInitials("John Doe", initials.WithSeparator(".")) // J.D
```

### WithSensitive

Keeps the case of the name.

```go
initials.GetInitials("John doe", initials.WithSensitive()) // Jd
```

### WithLowercase

Generates the initials in lowercase.

```go
initials.GetInitials("John Doe", initials.WithLowercase()) // jd
```

### WithUppercase

Generates the initials in uppercase.

```go
initials.GetInitials("John Doe", initials.WithUppercase()) // JD
```

### WithCamelCase

Capitalizes the first letter taken from each word and keeps the case of the others.

```go
initials.GetInitials("John Doe", initials.WithLength(3), initials.WithCamelCase()) // JDo
```

## Development

Prerequisites: [Task](https://taskfile.dev/installation/).

```shell
task go-tools:install    # install the tools at their pinned versions
task go:default          # clean, format, vet, tidy, lint and test
task go:test             # run the unit tests
task go:test-coverage    # generate the coverage report
```

## License

MIT, see [LICENSE](LICENSE).
