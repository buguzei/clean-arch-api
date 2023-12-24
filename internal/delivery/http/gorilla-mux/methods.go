package gorilla_mux

import (
	"log"
	"net/http"
)

func (gm GorillaMux) SignUp(w http.ResponseWriter, r *http.Request) {
	err := gm.uc.SignUp()
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
