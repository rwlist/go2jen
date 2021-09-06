package main

import (
	"encoding/json"
	"net/http"

	"github.com/rwlist/go2jen/pkg/jengen"
)

type Request struct {
	Code string
}

type Response struct {
	Code        string
	JenniferGen string
}

func process(req *Request) (*Response, error) {
	code := req.Code

	jg, err := jengen.Convert([]byte(code))
	if err != nil {
		return nil, err
	}

	return &Response{
		Code:        code,
		JenniferGen: string(jg),
	}, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		req Request
		res *Response
		err error
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err == nil {
		res, err = process(&req)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
