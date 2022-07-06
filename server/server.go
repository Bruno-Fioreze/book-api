package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.Server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(
		router.Run(":" + s.port),
	)
}
