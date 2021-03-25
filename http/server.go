package http

import (
	"net/http"
	"time"

	"ecnu.space/tmp-loser/conf"
	"github.com/gin-gonic/gin"
)

// Server serve http service
type Server struct{
	*http.Server
	g *gin.Engine
}

// NewServer init server
func NewServer(c *conf.AppConfig, opts ...Option)(s *Server){
	gin.SetMode(c.GinMode)
	s = &Server {g: gin.New()}
	s.Server = &http.Server{
		Addr:				c.ServerHost,
		Handler: 			s.g,
		ReadHeaderTimeout:	time.Second * 10,
		WriteTimeout:		time.Second * 10,
	}
	for _, opt := range opts {
		opt(s.g)
	}
	return 
}
