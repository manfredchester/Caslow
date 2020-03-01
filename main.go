package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			msg := trace("ERROR: %v", e)
			for _, m := range msg {
				fmt.Println(m)
			}
		}
	}()
	conf := flag.String("conf", "", "configuration file")
	ver := flag.Bool("version", false, "show version info")
	flag.Parse()

	if *ver {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "	")
		enc.Encode(version())
		return
	}

	parseConfig(*conf)
	savePid()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler(home))

	timeout := 60
	svr := http.Server{
		Addr:         ":" + rc.SERVICE_PORT,
		Handler:      mux,
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
	}
	assert(svr.ListenAndServe())
}

func savePid() {
	f, err := os.Create(rc.PID_FILE)
	assert(err)
	defer f.Close()
	_, err = f.Write([]byte(strconv.Itoa(os.Getpid())))
	assert(err)
}
