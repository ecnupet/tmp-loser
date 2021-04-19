package http

import (
	"net/http"
	"time"

	"ecnu.space/tmp-loser/conf"
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
	// g.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://www.ecnu.space", "https://*.ecnu.space", "http://localhost:3070", "http://discreater.asuscomm.com:3070"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
	// 	AllowHeaders:     []string{"Origin", "x-requested-with", "Referer", "User-Agent", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
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
