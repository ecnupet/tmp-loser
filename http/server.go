package http

import (
	"net/http"
	"time"

	"ecnu.space/tmp-loser/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server serve http service
type Server struct {
	*http.Server
	g *gin.Engine
}

// NewServer init server
func NewServer(c *conf.AppConfig, opts ...Option) (s *Server) {
	gin.SetMode(c.GinMode)
	g := gin.Default()
	// cors
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "x-requested-with", "Referer", "User-Agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	s = &Server{g: g}
	s.Server = &http.Server{
		Addr:              c.ServerHost,
		Handler:           s.g,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 10,
	}
	for _, opt := range opts {
		opt(s.g)
	}
	return
}
