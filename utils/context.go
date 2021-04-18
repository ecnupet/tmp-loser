package utils

import (
	"github.com/gin-gonic/gin"
)

type HttpState int
type Detail string

const (
	// http state
	OK HttpState = 200
	// 临时重定向
	Redirect      HttpState = 302
	Unknown       HttpState = 404
	BadRequest    HttpState = 400
	InternalError HttpState = 500

	//返回状态描述， 同步 C# person-manage服务的http返回格式，具体还得问hcs
	GetSuccess    Detail = "获取成功"
	GetFail		Detail = "获取失败"
	GetFailForDB  Detail = "获取失败，数据库错误"
	GetFailForNum Detail = "获取失败，题目不足"
	PostSuccess   Detail = "上传成功"
	PostFail      Detail = "上传失败, 参数错误"
	PostFailForDB Detail = "上传失败，数据库错误"
	DeleteSuccess Detail = "删除成功"
	DeleteFail    Detail = "删除失败"
)

type Response struct {
	State  HttpState   `json:"state"`
	Detail Detail      `json:"detail"`
	Data   interface{} `json:"data"`
}

func HandleGetDBErr(c *gin.Context, errString string) {
	c.JSON(int(InternalError), Response{
		State:  1,
		Detail: GetFailForDB,
		Data:   errString,
	})
}
func HandleGetErr(c *gin.Context, errString string) {
	c.JSON(int(BadRequest), Response{
		State:  4,
		Detail: GetFail,
		Data:   errString,
	})
}
func HandleGetNumErr(c *gin.Context, errString string) {
	c.JSON(int(InternalError), Response{
		State:  1,
		Detail: GetFailForNum,
		Data:   errString,

	})
}

func HandleGetSuccess(c *gin.Context, data interface{}) {
	c.JSON(int(OK), Response{
		State:  0,
		Detail: GetSuccess,
		Data:   data,
	})
}

func HandlePostSuccess(c *gin.Context, data interface{}) {
	c.JSON(int(OK), Response{
		State:  0,
		Detail: PostSuccess,
		Data:   data,
	})
}

func HandlePostQuizQuestionErr(c *gin.Context, errString string) {
	c.JSON(int(BadRequest), Response{
		State:  4,
		Detail: PostFail,
		Data:   errString,
	})
}

func HandlePostDBErr(c *gin.Context, errString string) {
	c.JSON(int(InternalError), Response{
		State:  1,
		Detail: PostFailForDB,
		Data:   errString,
	})
}

func GrpcErr(c *gin.Context, errString string) {
	c.JSON(401, struct {
		State  int    `json:"state"`
		Detail string `json:"detail"`
		Data   string `json:"data"`
	}{
		State:  7,
		Detail: "grpc请求失败",
		Data:   errString,
	})
}

func AuthErr(c *gin.Context, errString string) {
	c.JSON(401, struct {
		State  int    `json:"state"`
		Detail string `json:"detail"`
		Data   string `json:"data"`
	}{
		State:  7,
		Detail: "鉴权失败, 请登录",
		Data:   "",
	})
}

func ExtractCookieErr(c *gin.Context, errString string) {
	c.JSON(401, struct {
		State  int    `json:"state"`
		Detail string `json:"detail"`
		Data   string `json:"data"`
	}{
		State:  7,
		Detail: "提取Cookie失败",
		Data:   errString,
	})
}

func HandleNotAdminErr(c *gin.Context, errString string){
	c.JSON(401, struct {
		State  int    `json:"state"`
		Detail string `json:"detail"`
		Data   string `json:"data"`
	}{
		State:  7,
		Detail: "请确定您是否是管理员",
		Data:   errString,
	})
}

func HandleNullReturn(c *gin.Context, errString string) {
	c.JSON(404, struct {
		State  int    `json:"state"`
		Detail string `json:"detail"`
		Data   string `json:"data"`
	}{
		State:  404,
		Detail: "返回空",
		Data:   errString,
	})
}