package server

import (
	"github.com/jmoiron/sqlx"
	"todoapp/todoapp/repo"
	"todoapp/todoapp/service"
)

// InitServer initializes server
func InitServer(db *sqlx.DB) *Server {
	repo := repo.NewRepository(db)
	s := service.NewService(repo)
	return NewServer(s)
}
