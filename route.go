package main

import (
	"net/http"
	"net/url"
)

func home(args url.Values) interface{} {

	path := args.Get("REQUEST_URL_PATH")
	if path == "/" {

	}

	return httpError{
		Code: http.StatusNotFound,
		Mesg: "not found: " + path,
	}
}
