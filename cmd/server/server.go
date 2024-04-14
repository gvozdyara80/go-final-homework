package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-final-homework/handlers"
	"github.com/go-final-homework/repository"
	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(port string, db *sql.DB) *Server {
	return &Server{
		addr: ":" + port,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	
	userRepo := repository.NewRepository(s.db)
	userHandler := handlers.NewUserHandler(userRepo)
	userHandler.InitUserRoutes(router)

	transactionRepo := repository.NewRepository(s.db)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo)
	transactionHandler.InitTransactionRoutes(router)
	

	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, router)
}
