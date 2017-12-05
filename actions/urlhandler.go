package actions

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/clbanning/mxj"
	"github.com/go-chi/render"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
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

	render.JSON(w, r, mapVal)
}
