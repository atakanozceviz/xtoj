package actions

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/clbanning/mxj"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	xml, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer r.Body.Close()

	if len(xml) == 0 {
		io.WriteString(w, "post request body can not be empty")
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
}
