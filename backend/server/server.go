package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joptim/awesome-ost/backend/server/player"
	"net/http"
)

type Server struct {
	router *mux.Router
	player player.IPlayer
}

func (s *Server) ListenAndServe(addr string) error {
	if err := s.setUp(); err != nil {
		return err
	}
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) setUp() error {
	var err error
	s.player, err = player.New()
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	tracks := router.PathPrefix("/tracks").Subrouter()
	tracks.HandleFunc("/", s.tracksList).Methods("GET")
	tracks.HandleFunc("/remove/", s.tracksRemoveAll).Methods("PUT")
	tracks.HandleFunc("/{track}/add/", s.tracksAdd).Methods("PUT")
	tracks.HandleFunc("/{track}/remove/", s.tracksRemove).Methods("PUT")

	s.router = router
	return nil
}

func (s *Server) tracksList(w http.ResponseWriter, _ *http.Request) {
	// Todo: Read path from env variable or .env file
	assets, err := s.player.List()
	if err != nil {
		http.Error(w, "cannot load media folder", 500)
	}
	encoded, err := json.Marshal(assets)
	if err != nil {
		http.Error(w, "cannot encode asset list into json", 500)
	}
	if n, err := w.Write(encoded); n != len(encoded) || err != nil {
		http.Error(w, "unexpected error", 500)
	}
}

func (s *Server) tracksAdd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	track := vars["track"]
	if err := s.player.Add(track); err != nil {
		http.Error(w, fmt.Sprintf("cannot add track %s: %v", track, err), 404)
	}
	if err := s.player.Play(); err != nil {
		http.Error(w, fmt.Sprintf("cannot add track %s, %v", track, err), 500)
	}
	w.WriteHeader(200)
}

func (s *Server) tracksRemove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	track := vars["track"]

	if err := s.player.Remove(track); err != nil {
		http.Error(w, fmt.Sprintf("cannot remove track %s: %v", track, err), 404)
	}
	w.WriteHeader(200)
}

func (s *Server) tracksRemoveAll(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet", 400)
}

func New() *Server {
	return &Server{}
}
