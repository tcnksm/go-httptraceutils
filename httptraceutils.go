// Package httptraceutils is helper package for httptrace.
//
// Go 1.7 introduces net/http/httptrace package. That provides mechanisms for tracing within HTTP requests.
// Since it only provides the low level struct to set hooks for various states of request,
// you need to write logging or displaying part by yourself.
// go-httptraceutils is a small helper package for logging out each hook info of httptrace.
package httptraceutils

import (
	"context"
	"log"
	"net/http/httptrace"
	"os"
	"strings"
)

var defaultLogger = log.New(os.Stderr, "", log.LstdFlags)

// WithClientTrace wraps `httptrace.WithClientTrace` and
// sets logging hook in each step function of `httptrace.ClientTrace`.
func WithClientTrace(ctx context.Context) context.Context {
	return withClientTraceLogger(ctx, defaultLogger)
}

func withClientTraceLogger(ctx context.Context, logger *log.Logger) context.Context {
	return httptrace.WithClientTrace(ctx, newClientTrace(logger))
}

func newClientTrace(logger *log.Logger) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			logger.Printf("[GetConn] HostPort: %s", hostPort)
		},

		GotConn: func(info httptrace.GotConnInfo) {
			logger.Printf("[GotConn] LocalAddr: %s", info.Conn.LocalAddr())
			logger.Printf("[GotConn] RemoteAddr: %s", info.Conn.RemoteAddr())
			logger.Printf("[GotConn] Reused: %v", info.Reused)
			logger.Printf("[GotConn] WasIdle: %#v", info.WasIdle)
			if info.WasIdle {
				logger.Printf("[GotConn] IdleTime: %#v", info.IdleTime)
			}
		},

		PutIdleConn: func(err error) {
			if err != nil {
				logger.Printf("[PutIdeConn] Error: %s", err)
				return
			}
			logger.Printf("[PutIdeConn] Error: nil")
		},

		GotFirstResponseByte: func() {
			logger.Printf("[GotFirstResponseByte]")
		},

		Got100Continue: func() {
			logger.Printf("[Got100Continue]")
		},

		DNSStart: func(info httptrace.DNSStartInfo) {
			logger.Printf("[DNSStart] Host: %s", info.Host)
		},

		DNSDone: func(info httptrace.DNSDoneInfo) {
			addrs := make([]string, 0, len(info.Addrs))
			for _, addr := range info.Addrs {
				addrs = append(addrs, addr.String())
			}
			logger.Printf("[DNSDone] Addrs: %s", strings.Join(addrs, ","))
			logger.Printf("[DNSDone] Coalesced: %v", info.Coalesced)
			if info.Err != nil {
				logger.Printf("[DNSDone] Error: %s", info.Err)
				return
			}
			logger.Printf("[DNSDone] Error: nil")
		},

		ConnectStart: func(network, addr string) {
			logger.Printf("[ConnectStart] Network: %s", network)
			logger.Printf("[ConnectStart] Addr: %s", addr)
		},

		ConnectDone: func(network, addr string, err error) {
			logger.Printf("[ConnectDone] Network: %s", network)
			logger.Printf("[ConnectDone] Addr: %s", addr)
			if err != nil {
				logger.Printf("[ConnectDone] Error: %s", err)
				return
			}
			logger.Printf("[ConnectDone] Error: nil")
		},

		WroteHeaders: func() {
			logger.Printf("[WroteHeaders]")
		},

		Wait100Continue: func() {
			logger.Printf("[Wait100Continue]")
		},

		WroteRequest: func(info httptrace.WroteRequestInfo) {
			if info.Err != nil {
				logger.Printf("[WroteRequest] Error: %s", info.Err)
				return
			}
			logger.Printf("[WroteRequest] Error: nil")
		},
	}
}
