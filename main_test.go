package main

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestHandler(_ *testing.T) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`{"Code": "package a\n\nfunc hello(arg string) {\n\tfmt.Println(\"Hello\", arg)\t\t\n}"}`)

	req := httptest.NewRequest("POST", "/test", buf)
	w := httptest.NewRecorder()

	Handler(w, req)
	fmt.Println(w.Body)
}
