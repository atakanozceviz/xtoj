package actions

import (
	"io"
	"io/ioutil"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer r.Body.Close()

	if len(data) == 0 {
		io.WriteString(w, "post request body can not be empty")
		return
	}
	convert(data, w, r)
}
