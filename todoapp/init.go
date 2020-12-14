package todoapp

import (
	"github.com/jmoiron/sqlx"
	"todoapp/todoapp/repo"
	"todoapp/todoapp/server"
	"todoapp/todoapp/service"
)

// InitServer initializes server
func InitServer(db *sqlx.DB) *server.Server {
	repo := repo.NewRepository(db)
	s := service.NewService(repo)
	return server.NewServer(s)
}
