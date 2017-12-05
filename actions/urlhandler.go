package actions

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/clbanning/mxj"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := r.FormValue("url")
	if u == "" {
		io.WriteString(w, "url query string parameter can not be empty")
		return
	}

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
	defer resp.Body.Close()

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
}
