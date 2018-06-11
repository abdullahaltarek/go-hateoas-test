package main

import (
	"testing"
	"net/http"
)

func TestServer(t *testing.T) {
	res, _ := http.Get("http://127.0.0.1:8021/api/books")

	if res.StatusCode != 200 {
		t.Fatal("Serer not running")
	}
}