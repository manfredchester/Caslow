package main

import (
	"database/sql"
	"net/http"
	"net/url"
)

type (
	queryResult  map[string]interface{}
	queryResults []queryResult
	// keyInfo      struct {
	// key string
	// pos int
	// }
)

func query(args url.Values) (res interface{}) {
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
		err      error
	)
	dss = append(dss, ds)
	defer func() {
		if e := recover(); e != nil {
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
		// TODO ...
	}
	return nil
}

func doqry(conn *sql.DB, args url.Values) (queryResults, float64, float64) {

}
