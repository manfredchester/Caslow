package main

import (
	"fmt"
	"net/url"
)

type epArg struct {
	Name string
	Hint string
}

type endPoint struct {
	EndPoint string
	Purpose  string
	Note     string
	Args     []epArg
	Returns  []epArg
}

func help(args url.Values) interface{} {
	fmt.Println("help handle")
	path := args.Get("REQUEST_URL_PATH")
	fmt.Println("args.Get(REQUEST_URL_PATH)", path)
	return []endPoint{
		endPoint{
			EndPoint: "",
			Purpose:  "",
		},
		endPoint{
			EndPoint: "",
			Purpose:  "",
			Note:     "",
			Args: []epArg{
				epArg{
					Name: "",
					Hint: "",
				},
			},
			Returns: []epArg{
				epArg{
					Name: "",
					Hint: "",
				},
			},
		},
	}
}

func errpage(args url.Values) interface{} {
	return map[string]string{
		"ErrorInfo": args.Get("err"),
	}
}
