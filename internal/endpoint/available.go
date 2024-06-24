package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/nSakkriou/utils/pkg/agent"
)

func AvailableEndpoint(endCommandList []agent.EndCommand) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w = setHeader(w)

		if err := json.NewEncoder(w).Encode(endCommandList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
