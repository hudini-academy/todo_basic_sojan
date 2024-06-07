package main

import "net/http"

//func (app *application) routes() *http.ServeMux {
func (app *application) routes() http.Handler {
// http.Handler instead of *http.ServeMux.

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.List)
	mux.HandleFunc("/list", app.List)
	mux.HandleFunc("/tasks/add", app.Add)
	mux.HandleFunc("/tasks/listsingle", app.ListSingle)
	mux.HandleFunc("/tasks/delete", app.Delete)
	mux.HandleFunc("/tasks/update", app.Update)
	mux.HandleFunc("/tasks/updateform", app.UpdateForm)

	//file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

/// Pass the servemux as the 'next' parameter to the secureHeaders middleware
// Because secureHeaders is just a function, and the function returns a
// http.Handler we don't need to do anything else
return app.recoverPanic(app.logRequest(secureHeaders(mux)))

}
