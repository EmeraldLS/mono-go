# MonoGo

[![Go Reference](https://pkg.go.dev/badge/github.com/emeraldls/monogo.svg)](https://pkg.go.dev/github.com/emeraldls/monogo) [![Go Report Card](https://goreportcard.com/badge/github.com/emeraldls/monogo)](https://goreportcard.com/report/github.com/emeraldls/monogo)

- #### Still undergoing development, not ready for use.

<img align="right" alt="DiscordGo logo" src="docs/img/golang.png" width="300" height="350">

MonoGo is a [Go](https://golang.org/) package that provides low level
bindings to the [Mono](https://mono.co/) client API. MonoGo
has nearly complete support for all of the Mono API endpoints.

**For more info related to this package, visit the official [Mono](https://discord.gg/golang) documentation.**

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` _will always pull the latest tagged release from the master branch._

```sh
go get github.com/emeraldls/monogo
```

### Usage

Import the package into your project.

```go
import "github.com/emeraldls/monogo"
```

Construct a new Mono client which can be used to access the variety of
Mono API endpoints.

```go
mono := monogo.New("secret key")
```

See Documentation and Examples below for more detailed information.

## Documentation

**NOTICE**: This library and the Mono API are unfinished.
Because of that there may be major changes to library in the future.

The MonoGo code is fairly well documented at this point and is currently
the only documentation available. Go reference (below) presents that information in a nice format.

- [![Go Reference](https://pkg.go.dev/badge/github.com/emeraldls/monogo.svg)](https://pkg.go.dev/github.com/emeraldls/monogo)
- Hand crafted documentation coming eventually.

## Examples

## Troubleshooting

## Contributing

Contributions are very welcomed, however please follow the below guidelines.

- First open an issue describing the bug or enhancement so it can be
  discussed.
- Try to match current naming conventions as closely as possible.
- This package is intended to be a low level direct mapping of the Mono API,
  so please avoid adding enhancements outside of that scope without first
  discussing it.
- Create a Pull Request with your changes against the master branch.
