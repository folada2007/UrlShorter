package apiserver

import (
	"ShorterAPI/internal/app/apiserver/handler"
	"ShorterAPI/internal/domain/shorter"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	repo   shorter.Repository
	logger *logrus.Logger
	router *mux.Router
	config *Config
}

func New(cfg *Config, repo shorter.Repository) *APIServer {

	return &APIServer{
		repo:   repo,
		logger: logrus.New(),
		config: cfg,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Infof("Starting API Server in host %s", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {

	h := handler.NewHandler(s.logger, s.repo)
	s.router.HandleFunc("/", h.HomeHandler()).Methods("POST")
	s.router.HandleFunc("/{shortKey}", h.RedirectHandler()).Methods("GET")
}
