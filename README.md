# Generate initials from names

[![GoDoc](https://godoc.org/github.com/go-mods/initials?status.svg)](https://godoc.org/github.com/go-mods/initials)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mods/initials)](https://goreportcard.com/report/github.com/go-mods/initials)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/go-mods/initials/blob/master/LICENSE)

Initials is a Go package that generates initials from names.

## Installation

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
    fmt.Println(initials.GetInitials("John Doe"))                                   // Output: JD
    fmt.Println(initials.GetInitials("John Doe", WithLength(3), WithCamelCase()))   // Output: JDo
    fmt.Println(initials.GetInitials("John"))                                       // Output: JO  
    fmt.Println(initials.GetInitials("John"), WithCamelCase())                      // Output: Jo
}
```

## Options

### WithLength (_Default: 2_)
Defines the length of the initials

```go
initials.GetInitials("John Doe", WithLength(3)) // Output: JDO
```

## WithWordLength
Defines the length of the initials to the number of words

```go
initials.GetInitials("John David Dugga", WithWordLength()) // Output: JDD
```

## WithSeparator (_default: ""_)

Set the separator between the initials.

```go
initials.GetInitials("John Doe", WithSeparator(".")) // Output: J.D
```

## WithSensitive (_default: false_)

If set to true, the initials will be generated with the case of the words.

```go
initials.GetInitials("John doe", WithSensitive()) // Output: Jd
```

## WithLowercase (_default: false_)

If set to true, the initials will be generated in lowercase.

```go
initials.GetInitials("John Doe", WithLowercase()) // Output: jd
```

## WithUpperCase (_default: true_)

If set to true, the initials will be generated in uppercase.

```go
initials.GetInitials("John Doe", WithUpperCase()) // Output: JD
```

## WithCamelCase (_default: false_)

If set to true, the initials will be generated in camel case.

```go
initials.GetInitials("John Doe", WithLength(3), WithCamelCase()) // Output: JDo
```



