package httptraceutils

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestWithClientTrace(t *testing.T) {
}

func TestWithClientTraceLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.LstdFlags)

	req, err := http.NewRequest("GET", "http://localhost", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := withClientTraceLogger(context.Background(), logger)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	client.Do(req)

	if !strings.Contains(buf.String(), "[GetConn]") {
		t.Fatalf("expect [GetConn] is logged")
	}

	if !strings.Contains(buf.String(), "[DNSStart]") {
		t.Fatalf("expect [DNSStart] is logged")
	}
}
