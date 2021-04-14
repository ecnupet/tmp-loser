package http

import (
	"log"
	"net/http"

	"ecnu.space/tmp-loser/conf"
	"ecnu.space/tmp-loser/http/handler/api"
	"github.com/gin-gonic/gin"
	"ecnu.space/tmp-loser/http/handler/middleware"
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
	e.Use(middleware.Auth())
	authed := e.Group("/api/tl")
	// POST route
	authed.POST("test", api.TestPost)
	authed.POST("test2", api.TestPost2Seconds)
	authed.POST("/quiz/new", api.GenQuiz)
	authed.POST("/quiz/correct", api.CommitQuizQuestion)
	// 问题插入接口
	authed.POST("/question/insert", api.InsertQuestion)
	// GET route
	authed.GET("/question/detail", api.GetQuestionDetail)
	authed.GET("/quiz/history", api.GetQuizHistory)
	authed.GET("/quiz/history/detail", api.GetQuizHistoryDetail)
	authed.GET("/quiz/history/num", api.GetQuizHistoryNum)
	authed.GET("/statistics", api.GetStatistics)
	// 问题更新接口
}
