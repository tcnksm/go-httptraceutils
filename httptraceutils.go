// Package httptraceutils is helper tool for httptrace
package httptraceutils

import (
	"context"
	"log"
	"net/http/httptrace"
	"os"
)

var defaultLogger = log.New(os.Stderr, "", log.LstdFlags)

func WithClientTrace(ctx context.Context) context.Context {
	return WithClientTraceLogger(ctx, defaultLogger)
}

func WithClientTraceLogger(ctx context.Context, logger *log.Logger) context.Context {
	return httptrace.WithClientTrace(ctx, newClientTrace(logger))
}

func newClientTrace(logger *log.Logger) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			logger.Printf("[GetConn] hostPort: %s", hostPort)
		},

		GotConn: func(info httptrace.GotConnInfo) {
			logger.Printf("[GotConn] gotConnInfo: %#v", info)
		},

		PutIdleConn: func(err error) {
			logger.Printf("[PutIdeConn] err: %s", err)
		},

		GotFirstResponseByte: func() {
			logger.Printf("[GotFirstResponseByte]")
		},

		Got100Continue: func() {
			logger.Printf("[Got100Continue]")
		},

		DNSStart: func(info httptrace.DNSStartInfo) {
			// logger.Printf("[DNSStart] dnsStartInfo: %#v", info)
		},

		DNSDone: func(info httptrace.DNSDoneInfo) {
			//logger.Printf("[DNSDone] dnsDoneInfo: %#v", info)
		},

		ConnectStart: func(network, addr string) {
			logger.Printf("[ConnectStart] network: %s", network)
			logger.Printf("[ConnectStart] addr: %s", addr)
		},

		ConnectDone: func(network, addr string, err error) {
			logger.Printf("[ConnectDone] network: %s", network)
			logger.Printf("[ConnectDone] addr: %s", addr)
			if err != nil {
				logger.Printf("[ConnectDone] err: %s", err)
			}
		},

		WroteHeaders: func() {
			logger.Printf("[WroteHeaders]")
		},

		Wait100Continue: func() {
			logger.Printf("[Wait100Continue]")
		},

		WroteRequest: func(info httptrace.WroteRequestInfo) {
			logger.Printf("[WroteRequest]")
			if info.Err != nil {
				logger.Printf("[WroteRequest] err: %s", info.Err)
			}
		},
	}
}
