package server

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"todoapp/todoapp/client"
	"todoapp/todoapp/repo"
	"todoapp/todoapp/service"
)

// InitServer initializes server
func InitServer(db *sqlx.DB, conn *grpc.ClientConn) *Server {
	repoInstance := repo.NewRepository(db)
	clientInstance := client.NewEventClient(conn)
	s := service.NewService(repoInstance, clientInstance)
	return NewServer(s)
}
