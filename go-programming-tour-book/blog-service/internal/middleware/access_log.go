package middleware

import (
	"bytes"
	"time"

	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/global"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p) //向buffer写入后，再完成原有功能
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method: %s, status_code: %d, " + "begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(s,
			c.Request.Method,
			bodyWriter.ResponseWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
