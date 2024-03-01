package main

import "net/http"

func (a *application) get(s string, h http.HandlerFunc) {
	a.App.Routes.Get(s, h)
}

func (a *application) post(s string, h http.HandlerFunc) {
	a.App.Routes.Post(s, h)
}

func (a *application) put(s string, h http.HandlerFunc) {
	a.App.Routes.Put(s, h)
}

func (a *application) delete(s string, h http.HandlerFunc) {
	a.App.Routes.Delete(s, h)
}

func (a *application) patch(s string, h http.HandlerFunc) {
	a.App.Routes.Patch(s, h)
}

func (a *application) options(s string, h http.HandlerFunc) {
	a.App.Routes.Options(s, h)
}

func (a *application) head(s string, h http.HandlerFunc) {
	a.App.Routes.Head(s, h)
}

func (a *application) use(m ...func(http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}
