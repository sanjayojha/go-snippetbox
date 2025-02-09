package main

import (
	"net/http"

	"github.com/justinas/alice"
	"snippetbox.sanjayojha.in/ui"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// dynamic middleware
	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamicMiddleware.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamicMiddleware.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamicMiddleware.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamicMiddleware.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamicMiddleware.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamicMiddleware.ThenFunc(app.userLoginPost))

	// require authentication middleware
	protectedMiddleware := dynamicMiddleware.Append(app.requireAuthentication)
	mux.Handle("GET /snippet/create", protectedMiddleware.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protectedMiddleware.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protectedMiddleware.ThenFunc(app.userLogoutPost))

	//return app.recoverPanic(app.logRequest(commonHeaders(mux)))

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standardMiddleware.Then(mux)
}
