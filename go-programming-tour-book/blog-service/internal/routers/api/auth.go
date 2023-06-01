package api

import (
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/global"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/internal/service"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/app"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindandValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param) //验证数据库中拥有以后，就可以建立token了
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
