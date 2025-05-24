package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"packr/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api", func(api chi.Router) {
		api.Use(s.AuthMiddleware)

		api.Get("/packs", s.GetPacksHandler)
		api.Post("/packs", s.PostPacksHandler)
		api.Delete("/packs", s.DeletePacksHandler)

		api.Get("/solution", s.GetSolutionHandler)
	})

	store.ImportPacks(os.Getenv("INIT_PACKS"))

	return r
}

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("X-Api-Key") != s.apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) GetPacksHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Packs []int `json:"packs"`
	}{
		Packs: store.GetPacks(),
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) GetSolutionHandler(w http.ResponseWriter, r *http.Request) {
	amountStr := r.URL.Query().Get("amount")

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := struct {
		Solution map[int]int `json:"solution"`
	}{
		Solution: store.Solve(amount),
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) PostPacksHandler(w http.ResponseWriter, r *http.Request) {
	packStr := r.URL.Query().Get("pack")

	pack, err := strconv.Atoi(packStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.AddPack(pack)

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) DeletePacksHandler(w http.ResponseWriter, r *http.Request) {
	packStr := r.URL.Query().Get("pack")

	pack, err := strconv.Atoi(packStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.RemovePack(pack)

	w.WriteHeader(http.StatusOK)
}
