// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func(app *application) routes() *httprouter.Router {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
  router.Handler(http.MethodGet,"/static/*filepath", http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet,"/", app.home)
	router.HandlerFunc(http.MethodGet, "/about", app.about)
	router.HandlerFunc(http.MethodGet,"/poll/reply", app.pollReplyShow)
	router.HandlerFunc(http.MethodPost, "/poll/reply", app.pollReplySubmit)
	router.HandlerFunc(http.MethodGet, "/poll/success", app.pollSuccessShow)
	router.HandlerFunc(http.MethodGet, "/poll/create", app.pollCreateShow)
	router.HandlerFunc(http.MethodPost, "/poll/create", app.pollCreateSubmit)
	router.HandlerFunc(http.MethodGet, "/options/create", app.optionsCreateShow)
	router.HandlerFunc(http.MethodPost, "/options/create", app.optionsCreateSubmit)
	
	return router
}