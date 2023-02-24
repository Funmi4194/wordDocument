package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	response "github.com/Funmi4194/myMod/responses"
	"github.com/Funmi4194/myMod/types"
)

// BodyWithFunc is a middleware that parses the request body into the given interface
// and stores it in the request context.
func BodyWithFunc(i interface{}) func(func(http.ResponseWriter, *http.Request)) http.Handler {
	return func(next func(http.ResponseWriter, *http.Request)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//parse body

			err := json.NewDecoder(r.Body).Decode(i)
			if err != nil {
				if err.Error() != "EOF" {
					response.SendJSONResponse(w, false, http.StatusBadRequest, "Error parsing body: "+err.Error(), nil)
					return
				}
			}

			// set body in context
			ctx := context.WithValue(r.Context(), types.BodyCtxKey{}, i)
			next(w, r.WithContext(ctx))
		})
	}
}
