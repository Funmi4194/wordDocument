package middleware

import (
	"context"
	"net/http"

	"github.com/Funmi4194/myMod/types"
	"github.com/gorilla/mux"
)

//ParasWithFunc middleware  parses a variable params into the given request context.
func ParamsWithFunc() func(func(http.ResponseWriter, *http.Request)) http.Handler {
	return func(next func(http.ResponseWriter, *http.Request)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// set variable params
			params := mux.Vars(r)

			// Add the document params to the request context
			ctx := context.WithValue(r.Context(), types.ParamsCtxKey{}, params)

			next(w, r.WithContext(ctx))

		})
	}
}

func DocsWithFunc() func(func(http.ResponseWriter, *http.Request)) http.Handler {
	return func(next func(http.ResponseWriter, *http.Request)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := context.WithValue(r.Context(), types.AuthCtxKey{}, "document")

			next(w, r.WithContext(ctx))

		})
	}
}
