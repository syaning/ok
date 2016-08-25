# ok

The most simple but handy HTTP request library for Go.

## Installation

```sh
$ go get github.com/syaning/ok
```

## Getting Started

### HTTP Get

```go
str, err := ok.
    Get("http://httpbin.org/get").
    OK().
    ToString()
fmt.Println(str, err)
```

### Download

```go
size, err := ok.Download("http://httpbin.org/image/png", "img.png")
fmt.Println(size, err)
```

## License

[MIT](./LICENSE)