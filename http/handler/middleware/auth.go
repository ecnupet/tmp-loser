package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"ecnu.space/tmp-loser/model"
	"ecnu.space/tmp-loser/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("haha")
		token, err := c.Cookie(".AspNetCore.Cookies")
		if err != nil {
			utils.ExtractCookieErr(c, err.Error())
			log.Println(err.Error())
			c.Abort()
		}
		client := &http.Client{}
		//生成要访问的url
		url := "http://pm-app-svc.backend:5000/api/pm/auth/check"
		//提交请求
		req, err := http.NewRequest("GET", url, nil)
		//增加header选项
		req.Header.Add("Cookie", ".AspNetCore.Cookies="+token)
		req.Header.Add("Host", "www.ecnu.space")

		req.Header.Add("X-Requested-With", "XMLHttpRequest")
		if err != nil {
			utils.AuthErr(c, err.Error())
			log.Println("1", err.Error())
			c.Abort()
		}
		//处理返回结果
		resp, err := client.Do(req)
		if err != nil {
			utils.AuthErr(c, err.Error())
			log.Println("2", err.Error())
			c.Abort()
		}
		defer resp.Body.Close()
		auth := model.Auth{}
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(resp.StatusCode)
		fmt.Println("json: ", string(body))
		if err != nil {
			utils.AuthErr(c, err.Error())
			log.Println("3", err.Error())
			c.Abort()
		}
		err = json.Unmarshal(body, &auth)
		if err != nil {
			utils.AuthErr(c, err.Error())
			log.Println("4", err.Error())
			c.Abort()
		}

		c.Set("user_name", auth.Data.Name)
		c.Set("isAdmin", auth.Data.IsAdmin)
		// client := personManage.GetPMClient()

		// md := metadata.Pairs("Cookie", ".AspNetCore.Cookies="+token)
		// ctx := metadata.NewOutgoingContext(context.Background(), md)
		// header := metadata.New(map[string]string{"content-type": "application/grpc"})
		// err = grpc.SetHeader(ctx, header)
		// if err != nil {
		// 	log.Println("setHeader err " + err.Error())
		// }

		// md2 := metadata.Pairs("Content-Type", "application/grpc")
		// ctx = metadata.NewOutgoingContext(context.Background(), md2)
		// in := proto.HelloRequest{Name: "123"}
		// resp, err := client.Check(ctx, &in)
		// if err != nil {
		// 	util.GrpcErr(c, err.Error())
		// 	log.Println(err.Error())
		// 	c.Abort()
		// }
		// log.Println(resp.GetName())
		// c.Set("user_name", resp.GetName())
		c.Next()
	}
}
