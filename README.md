# ok

The most simple but handy HTTP request library for Go.

## Installation

```sh
$ go get github.com/syaning/ok
```

## Getting Started

The exported struct `ok.RequestWrapper` wraps a client, an HTTP request, an HTTP response and a potential error together. When sending a request and get respones, there are 4 steps:

1. Create a `RequestWrapper`
    - `Request()`
    - `NewRequest(method, urlStr string)`
    - `Get(urlStr string)`
    - `Post(urlStr string)`
    - `Put(urlStr string)`
    - `Delete(urlStr string)`
    - ...
2. Prepare the request: set headers, set body, use proxy, ...
    - `Method(method string)`
    - `Url(urlStr string)`
    - `Query(query string)`
    - `Set(key, value string)`
    - `Proxy(proxy string)`
    - ...
3. Send request
    - `OK()`
4. Get response and read response body
    - `ToBytes()`
    - `ToString()`
    - `Pipe(w io.Writer)`
    - `ToFile(filename string)`
    - ...

See full functions and methods on [GoDoc](https://godoc.org/github.com/syaning/ok).

## Examples

### HTTP Get

```go
str, err := ok.
    Get("http://httpbin.org/get").
    OK().
    ToString()
fmt.Println(str, err)
```

### HTTP Post JSON

```go
size, err := ok.
    Post("http://httpbin.org/post").
    Json(`{"greeting":"hello world"}`).
    OK().
    Pipe(os.Stdout)
fmt.Println(size, err)
```

### HTTP Post Form

```go
size, err := ok.
    Post("http://httpbin.org/post").
    Form("greeting=hello world").
    OK().
    ToFile("res.json")
fmt.Println(size, err)
```

### Download

```go
size, err := ok.Download("http://httpbin.org/image/png", "img.png")
fmt.Println(size, err)
```

## License

[MIT](./LICENSE)