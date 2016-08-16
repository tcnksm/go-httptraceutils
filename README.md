# go-httptraceutils

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[license]: https://github.com/tcnksm/go-httptraceutils/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/tcnksm/go-httptraceutils

[Go 1.7](https://tip.golang.org/doc/go1.7) introduces [`net/http/httptrace`](https://tip.golang.org/pkg/net/http/httptrace/) package. That provides mechanisms for tracing within HTTP requests. Since it only provides the low level struct to set hooks for various states of request, you need to write logging or displaying part by yourself. `go-httptraceutils` is a small helper package for logging out each hook info of `httptrace`. See [example](/_example). 

## Install 

Use `go get`,

```bash
$ go get github.com/tcnksm/go-httptraceutils
```

## Author

[Taichi Nakashima](https://github.com/tcnksm)
