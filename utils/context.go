package utils

import (
	"github.com/gin-gonic/gin"
)

type HttpState int
type Detail string

const (
	// http state
	OK            HttpState = 200
	// 临时重定向
	Redirect      HttpState = 302
	Unknown       HttpState = 404
	BadRequest    HttpState = 400
	InternalError HttpState = 500

	//返回状态描述， 同步 C# person-manage服务的http返回格式，具体还得问hcs
	GetSuccess    Detail = "获取成功"
	GetFailForDB  Detail = "获取失败，数据库错误"
	GetFailForNum Detail = "获取失败，题目不足"
	PostSuccess   Detail = "上传成功"
	PostFail      Detail = "上传失败, 参数错误"
	DeleteSuccess Detail = "删除成功"
	DeleteFail    Detail = "删除失败"
)

type Response struct {
	State  HttpState       `json:"state"`
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

func HandlePostQuizQuestion(c *gin.Context){
	c.JSON(int(BadRequest), Response{
		State: BadRequest,
		Detail: PostFail,
		Data: "",
	})
}