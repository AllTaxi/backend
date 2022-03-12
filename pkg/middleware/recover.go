package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

//PanicRecovery recovers from panic and returns 500 as errCode
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()), "Error: ",err)
			}
		}()
		next.ServeHTTP(w, req)
	})
}
