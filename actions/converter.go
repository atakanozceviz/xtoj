package actions

import (
	"io"
	"net/http"

	"github.com/clbanning/mxj"
	"github.com/go-chi/render"
)

func convert(data []byte, w http.ResponseWriter, r *http.Request) {
	// try xml to json
	xmlMap, xmlErr := mxj.NewMapXml(data, true)
	if xmlErr != nil {
		// try json to xml
		jsonMap, jsonErr := mxj.NewMapJson(data)
		if jsonErr != nil {
			io.WriteString(w, xmlErr.Error()+"\n"+jsonErr.Error())
			return
		}
		xmlVal, err := jsonMap.Xml()
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml")
		w.Write(xmlVal)
		return
	}
	render.JSON(w, r, &xmlMap)
}
