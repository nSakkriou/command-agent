package endpoint

import (
	"net/http"
)

func HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w = setHeader(w)

	output(w,
		struct {
			message string
		}{
			message: "Server is healthy",
		})
}
