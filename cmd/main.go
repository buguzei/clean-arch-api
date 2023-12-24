package main

import (
	clean_arch_api "clean-arch-api"
	gorilla_mux "clean-arch-api/internal/delivery/http/gorilla-mux"
	"clean-arch-api/internal/repo/postgres"
	"clean-arch-api/internal/usecase"
	"log"
)

func main() {
	server := clean_arch_api.HTTPServer{}

	pg := postgres.NewPostgres()

	uc := usecase.NewUseCase(pg)

	gm := gorilla_mux.NewGorillaMux(uc)

	if err := server.Run("8090", gm.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
