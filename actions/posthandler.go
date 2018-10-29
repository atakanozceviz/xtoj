package actions

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.PlainText(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	if len(data) == 0 {
		render.PlainText(w, r, "post request body can not be empty")
		return
	}
	convert(data, w, r)
}
