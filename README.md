# go-httptraceutils

[Go 1.7](https://tip.golang.org/doc/go1.7) introduces [`net/http/httptrace`](https://tip.golang.org/pkg/net/http/httptrace/) package. That provides mechanisms for tracing within HTTP requests. Since it only provides the low level struct to set hooks for various states of request, you need to write logging or displaying part by yourself. `go-httptraceutils` is small wrapper of `httptrace`. 
