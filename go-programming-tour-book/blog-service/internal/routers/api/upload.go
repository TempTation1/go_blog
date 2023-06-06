package api

import (
	"fmt"

	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/global"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/internal/service"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/app"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	//file, fileHeader, err := c.Request.FormFile("file") //FormFile读取form中的文件参数，返回就是multipul类的文件接口和指针，文件同步被读到了内存中
	var fileAccessUrlList []string
	form, err := c.MultipartForm()
	fmt.Println(form)
	fileHeaderList := form.File["file"]
	for _, fileHeader := range fileHeaderList {
		fileType := convert.SrcTo(c.PostForm("type")).MustInt() //没做参数校验，直接用postform获取了form中的key对应的value
		if err != nil {
			errRsp := errcode.InvalidParams.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		if fileHeader == nil || fileType <= 0 {
			response.ToErrorResponse(errcode.InvalidParams)
			return
		}
		//c.SaveUploadedFile()
		svc := service.New(c.Request.Context())
		file, _ := fileHeader.Open()
		fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
		if err != nil {
			global.Logger.Errorf("svc.UploadFile err: %v", err)
			errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		fileAccessUrlList = append(fileAccessUrlList, fileInfo.AccessUrl)
	}

	response.ToResponse(gin.H{
		"file_access_urls": fileAccessUrlList,
	})
}
