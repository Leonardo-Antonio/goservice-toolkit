package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Leonardo-Antonio/goservice-toolkit/response"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recover from:", err)
				response.Writer(w, response.Info{
					Status:  http.StatusInternalServerError,
					Message: "Internal Server Error",
					Data:    nil,
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
