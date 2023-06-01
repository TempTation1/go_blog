package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/global"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

//只加密前面部分，另外剥离出后缀
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMd5(fileName)

	return fileName + ext
}

//获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

//获取文件保存地址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//检查保存目录是否存在，对os.Stat返回的error值校验查看是不是对应的错误
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//查看文件后缀是不是支持
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

//查看文件大小是不是超出限制，超过了是true
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

//查看文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//创建存储目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm) //自带不存在时创建，存在时忽略的功能
	if err != nil {
		return err
	}

	return nil
}

//保存上传的文件，os创建对应路径的文件描述符，然后用copy进行io操作
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open() //有header就行，file只是fileheader.open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
