package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		fmt.Fprintf(w, "GET\n")
	case "PUT":
		fmt.Fprintf(w, "PUT\n")
	}

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(reqDump))
}

func main() {
	fmt.Println("quick")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
