package main

import (
	conf "comment/config"
	hd "comment/internals/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", hd.MainHandler)
	// log.Printf("Server (commentAPI) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}
