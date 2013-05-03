package rollatoon

import (
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", wowHandler).
		Methods("GET")
	r.HandleFunc("/horde", hordeHandler).
		Methods("GET")
	r.HandleFunc("/alliance", allianceHandler).
		Methods("GET")
	r.HandleFunc("/about", aboutHandler).
		Methods("GET")

	http.Handle("/", r)
}
