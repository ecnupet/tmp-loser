package http

import "github.com/gin-gonic/gin"

// Option 给gin提供handler
type Option func(g *gin.Engine)

// WithGinRouteOption add routes for gin
func WithGinRouteOption(route func(*gin.Engine)) Option {
	return func(g *gin.Engine){
		route(g)
	}
}