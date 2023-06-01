package middleware

import (
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/app"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		//每一个请求，先从url或者header尝试获取token
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" { //没token就不给访问
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token) //有token就做一下解析鉴权
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
