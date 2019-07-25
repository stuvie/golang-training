package main

/*
samples:
go get github.com/julienschmidt/httprouter
curl -X PUT localhost:8080/entry/first/hello
curl -X PUT localhost:8080/entry/second/hi
curl -X PUT localhost:8080/entry/third/foo
curl localhost:8080/list
curl localhost:8080/entry/first
*/

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	data map[string]string
)

func main() {
	flag.Parse()
	data = map[string]string{}
	r := httprouter.New()
	r.GET("/entry/:key", show)
	r.GET("/list", show)
	r.PUT("/entry/:key/:value", update)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")
	if k == "" {
		fmt.Fprintf(w, "read list: %v", data)
		return
	}
	fmt.Fprintf(w, "read entry: data[%s] = %s", k, data[k])
}

func update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")
	v := p.ByName("value")

	data[k] = v

	fmt.Fprintf(w, "updated: data[%s] = %s", k, data[k])
}
