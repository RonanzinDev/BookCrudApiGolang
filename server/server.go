package server

import (
	"log"

	"github.com/gin-gonic/gin"
	router "github.com/ronanzindev/book-api-golang/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.server)
	log.Printf("Server running at port : %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
