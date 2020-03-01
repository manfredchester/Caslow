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
	fmt.Println("help")

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
