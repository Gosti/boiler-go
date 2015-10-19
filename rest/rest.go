package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Have to find how to PUT (interface{} conversion)

func handler(w http.ResponseWriter, r *http.Request, list map[string]interface{}) {
	if r.Method == "POST" {
		r.ParseMultipartForm(0)
		fmt.Fprintf(w, "%#v", r.Form)
		for v, i := range r.Form {
			list[v] = i
		}
	} else if r.Method == "DELETE" {
		r.ParseMultipartForm(0)
		for v, _ := range r.Form {
			delete(list, v)
		}
	} else {
		v, _ := json.Marshal(list[r.URL.Path[1:]])
		fmt.Fprintf(w, "%s", v)
	}
}

func main() {
	var list map[string]interface{} = make(map[string]interface{})
	list["totot"] = "oll"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, list) // Use a Wrapper to make it easier to use the list
	})
	http.ListenAndServe(":8080", nil)
}
