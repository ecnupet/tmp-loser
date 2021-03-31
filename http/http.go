package http

import (
	"log"
	"net/http"

	"ecnu.space/tmp-loser/conf"
	"ecnu.space/tmp-loser/http/handler/api"
	"github.com/gin-gonic/gin"
)

// InitAndStart init http and start service
func InitAndStart(c *conf.AppConfig) *http.Server {
	s := NewServer(
		c,
		WithGinRouteOption(route),
	)
	go func() {
		log.Printf("running tmp-loser HTTP server on %s", c.ServerHost)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal("start server fail. errror: ", err)
		}
	}()
	return s.Server
}

//route provide http routes to invoke handler
func route(e *gin.Engine) {
	// group
	unauthed := e.Group("/api/tl")
	// route
	unauthed.POST("/quiz/new", api.GenQuiz)
	unauthed.POST("/quiz/correct", api.CommitQuizQuestion)
	unauthed.GET("/question/detail", api.GetQuestionDetail)
	unauthed.GET("/quiz/history", api.GetQuizHistory)
	unauthed.GET("/quiz/history/detail", api.GetQuizHistoryDetail)
}
