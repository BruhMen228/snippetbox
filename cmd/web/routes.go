package main

import "net/http"

func (a *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(`./ui/static`)})
	//mux.Handle("/static", http.NotFoundHandler())
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)
	
	return mux
}