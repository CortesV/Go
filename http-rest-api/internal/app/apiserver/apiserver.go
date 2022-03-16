package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}

	server.configureRouter()

	server.logger.Info("Starting API server ...")

	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	server.logger.SetLevel(level)

	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello")
	}
}
