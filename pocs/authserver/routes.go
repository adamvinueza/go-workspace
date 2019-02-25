package authserver

import (
	"github.com/adamvinueza/go-workspace/pocs/bcryptclient"
)

func validateFunction(w http.ResponseWriter, r *http.Request, method string) error {
	var err error
	if r.method != method {
		err = http.Error(w, fmt.Sprintf("method %v not allowed"), http.StatusMethodNotAllowed)
	}
	return err
}

func GetHandler(w http.ResponseWriter, req *http.Request) {
	if r.method != method {
		err = http.Error(w, fmt.Sprintf("method %v not allowed"), http.StatusMethodNotAllowed)
		io.WriteString(
	}
	w.Header().Set("Content-Type", "application/json")
}

func Get(w http.ResponseWriter, r *http.Request) {
	validateFunction(w, r, "GET")

}
