# gocovergate

`gocovergate` is a simple tool that acts as a code coverage quality gate.

It expects a `cover.out` file to be present in the current directory.

When executed, it will print the total code coverage.
If it is below 80%, it will exit with a status code of 1.

## Installation

```shell
go install github.com/patrickhoefler/gocovergate@latest
```

## Usage

```shell
go test ./... --coverprofile cover.out
gocovergate
```

## License
[MIT](LICENSE)
