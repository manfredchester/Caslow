package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func home(args url.Values) interface{} {
	fmt.Println("home")
	path := args.Get("REQUEST_URL_PATH")
	if path == "/" {
		return "please access /help"
	}
	return httpError{
		Code: http.StatusNotFound,
		Mesg: "not found: " + path,
	}
}
