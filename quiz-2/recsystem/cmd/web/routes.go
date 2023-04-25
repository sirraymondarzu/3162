// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave)
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/about", dynamicMiddleware.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/login", dynamicMiddleware.ThenFunc(app.loginform))
	router.Handler(http.MethodPost, "/login", dynamicMiddleware.ThenFunc(app.loginform))
	router.Handler(http.MethodGet, "/register", dynamicMiddleware.ThenFunc(app.register))
	router.Handler(http.MethodGet, "/reservation", dynamicMiddleware.ThenFunc(app.reserve))
	router.Handler(http.MethodGet, "/feedback", dynamicMiddleware.ThenFunc(app.feedback))
	router.Handler(http.MethodPost, "/user", dynamicMiddleware.ThenFunc(app.userPortal))
	router.Handler(http.MethodPost, "/admin", dynamicMiddleware.ThenFunc(app.adminPortal))
	
	//tidy up the middleware chain 
	standardMiddleware := alice.New(
		app.recoverPanicMiddleware, 
		app.logRequestMiddleware, 
		securityHeadersMiddleware,
	)

	return standardMiddleware.Then(router)
}
