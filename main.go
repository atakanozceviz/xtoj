package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/clbanning/mxj"
	"github.com/unrolled/secure"
)

func main() {
	PORT := os.Getenv("PORT")
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect:     true,
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})

	app := secureMiddleware.Handler(xtoj)
	http.ListenAndServe(":"+PORT, app)
}

var xtoj = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
