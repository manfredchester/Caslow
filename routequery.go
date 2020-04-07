package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"test/zhlog"
	"time"
)

type (
	queryResult  map[string]interface{}
	queryResults []queryResult
	// keyInfo      struct {
	// key string
	// pos int
	// }
)

func query(args url.Values,reqBody map[string]interface{}) (res interface{}) {
	use := args.Get("use")
	dl.RLock()
	ds, ok := dsns[use]
	dl.RUnlock()
	if !ok {
		return httpError{
			Code: http.StatusSeeOther,
			Mesg: "[use] is not a valid data source",
		}
	}
	var (
		dss      []dsInfo
		recs     queryResults
		tqs, tfs float64
	)
	dss = append(dss, ds)
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error("traceID string", "%s", e.(error).Error())
			res = httpError{
				Code: http.StatusInternalServerError,
				Mesg: e.(error).Error(),
			}
		}
	}()
	for _, ds := range dss {
		conn, err := getDB(ds.Driver, ds.Dsn, "query")
		assert(err)
		data, tq, tf := doqry(conn, args)
		tqs = tqs + tq
		tfs = tfs + tf
		for _, d := range data {
			if len(dss) > 1 {
				d[rc.DB_TAG] = ds.Name
			}
			recs = append(recs, d)
		}
		summary := fmt.Sprintf("Got %d row(s) in %fs (query=%fs; fetch=%fs)",
			len(recs), tqs+tfs, tqs, tfs)
		fmt.Println("summary:", summary)
		recs = append(recs, map[string]interface{}{
			"summary": summary,
		})
		// args.Set("RESTIQUE_SUMMARY", summary)
	}
	return recs
}

func doqry(conn *sql.DB, args url.Values) (queryResults, float64, float64) {
	var tq, tf float64
	qry := args.Get("sql")
	timeout := time.Duration(rc.QUERY_TIMEOUT) * time.Second
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	defer cf()
	start := time.Now()
	rows, err := conn.QueryContext(ctx, qry)
	assert(err)
	tq = time.Since(start).Seconds()
	start = time.Now()
	cols, err := rows.Columns()
	assert(err)
	raw := make([][]byte, len(cols))
	ptr := make([]interface{}, len(cols))
	for i := range raw {
		ptr[i] = &raw[i]
	}
	recs := queryResults{}
	RangeRows(rows, func() {
		assert(rows.Scan(ptr...))
		rec := queryResult{}
		for i, r := range raw {
			if r == nil {
				rec[cols[i]] = nil
			} else {
				rec[cols[i]] = string(r)
			}
		}
		if rc.QUERY_MAXROWS > 0 && len(recs) > rc.QUERY_MAXROWS {
			args.Set("RESTIQUE_MAXROW", "1")
			return
		}
		recs = append(recs, rec)
	})
	tf = time.Since(start).Seconds()
	return recs, tq, tf
}
