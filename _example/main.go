package main

import (
	"context"
	"log"
	"net/http"

	"github.com/tcnksm/go-httptraceutils"
)

func main() {
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := httptraceutils.WithClientTrace(context.Background())
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	_ = res
}
