package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	r.HandleFunc("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm))).Methods("POST")
	r.HandleFunc("/films", h.FindFilms).Methods("GET")
	r.HandleFunc("/film/{id}", middleware.Auth(h.GetFilm)).Methods("GET")
	r.HandleFunc("/film/{id}", middleware.Auth(h.UpdateFilm)).Methods("PATCH")
	r.HandleFunc("/film/{id}", middleware.Auth(h.DeleteFilm)).Methods("DELETE")
}
