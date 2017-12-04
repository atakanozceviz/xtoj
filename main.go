package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/clbanning/mxj"
)

func main() {
	PORT := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := r.FormValue("url")
		resp, err := http.Get(u)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		xml, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		mapVal, err := mxj.NewMapXml(xml, true)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		js, err := mapVal.Json()
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		w.Write(js)
	})
	http.ListenAndServe(":"+PORT, nil)
}
