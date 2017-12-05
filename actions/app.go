package actions

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/secure"
)

func Start(port string) error {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect:     true,
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
	r := chi.NewRouter()
	r.Use(secureMiddleware.Handler)
	r.Get("/", urlHandler)
	r.Post("/", postHandler)

	return http.ListenAndServe(":"+port, r)
}
