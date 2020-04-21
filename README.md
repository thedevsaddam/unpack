Unpack
---
[![Build Status](https://travis-ci.org/thedevsaddam/unpack.svg?branch=master)](https://travis-ci.org/thedevsaddam/unpack)
[![Project status](https://img.shields.io/badge/version-v1-green.svg)](https://github.com/thedevsaddam/unpack/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/thedevsaddam/unpack)](https://goreportcard.com/report/github.com/thedevsaddam/unpack)
[![Coverage Status](https://coveralls.io/repos/github/thedevsaddam/unpack/badge.svg?branch=master)](https://coveralls.io/github/thedevsaddam/unpack)
[![GoDoc](https://godoc.org/github.com/thedevsaddam/unpack?status.svg)](https://pkg.go.dev/github.com/thedevsaddam/unpack)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](LICENSE.md)

Go assignment by slice, array unpacking or destructuring

### Installation

Install the package using
```go
$ go get github.com/thedevsaddam/unpack
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/thedevsaddam/unpack"
```

Let's see a quick example:

```go
package main

import (
	"fmt"

	"github.com/thedevsaddam/unpack"
)

func main() {
	names := []string{"Captain Jack Sparrow", "Tom", "Jerry"}
	var jack, tom string
	unpack.Do(names, &jack, &tom)
	fmt.Println("I'm", jack)
}
```


## Bugs and Issues

If you encounter any bugs or issues, feel free to [open an issue at
github](https://github.com/thedevsaddam/unpack/issues).

Also, you can shoot me an email to
<mailto:thedevsaddam@gmail.com> for hugs or bugs.

## Contribution
If you are interested to make the package better please send pull requests or create an issue so that others can fix.
[Read the contribution guide here](CONTRIBUTING.md)


## License
The **unpack** is an open-source software licensed under the [MIT License](LICENSE.md).