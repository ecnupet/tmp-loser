package utils

import (
	"github.com/gin-gonic/gin"
)

type State int
type Detail string

const (
	// http state
	OK            State = 200
	Redirect      State = 302
	Unknown       State = 404
	BadRequest    State = 400
	InternalError State = 500

	GetSuccess    Detail = "获取成功"
	GetFailForDB  Detail = "获取失败，数据库错误"
	GetFailForNum Detail = "获取失败，题目不足"
	PostSuccess   Detail = "上传成功"
	PostFail      Detail = "上传失败"
	DeleteSuccess Detail = "删除成功"
	DeleteFail    Detail = "删除失败"
)

type Response struct {
	State  State       `json:"state"`
	Detail Detail      `json:"detail"`
	Data   interface{} `json:"data"`
}

func HandleGetDBErr(c *gin.Context) {
	c.JSON(int(InternalError), Response{
		State:  InternalError,
		Detail: GetFailForDB,
		Data:   []string{},
	})
}

func HandleGetNumErr(c *gin.Context) {
	c.JSON(int(InternalError), Response{
		State: InternalError,
		Detail: GetFailForNum,
		Data: []string{},
	})
}

func HandleGetSuccess(c *gin.Context, data interface{}) {
	c.JSON(int(OK), Response{
		State: OK,
		Detail: GetSuccess,
		Data: data,
	})
}