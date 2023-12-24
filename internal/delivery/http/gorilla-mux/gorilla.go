package gorilla_mux

import (
	"clean-arch-api/internal/usecase"
	"github.com/gorilla/mux"
)

type GorillaMux struct {
	gm *mux.Router
	uc usecase.UC
}

func NewGorillaMux(uc usecase.UC) GorillaMux {
	return GorillaMux{gm: mux.NewRouter(), uc: uc}
}

func (gm GorillaMux) InitRoutes() *mux.Router {
	gm.gm.HandleFunc("/sign-up", gm.SignUp)
	gm.gm.HandleFunc("/sign-in", nil)
	gm.gm.HandleFunc("/delete-account", nil)

	return gm.gm
}
