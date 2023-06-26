package upload

import (
	"fmt"
	"io"
	"time"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/util/xfile"
	"github.com/binbinly/pkg/util/xhash"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/ecode"
	"gin-chat/pkg/app"
)

const FileMaxSize = 10 << 20 //最大上传10MB

// File 上传文件
// @Summary 上传文件
// @Description 上传文件
// @Tags 上传
// @Accept		json
// @Produce		json
// @Param Token header string true "用户令牌"
// @Param file formData file true "文件"
// @success 0 {object} app.Response{data=string} "调用成功结构"
// @Router /upload/file [post]
func File(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	//限制大小
	if file.Size > FileMaxSize {
		app.Error(c, ecode.ErrUploadImageLimit)
	}
	//打开文件
	f, err := file.Open()
	if err != nil {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	//文件名取md5值
	filename, err := xhash.MD5Reader(f)
	if err != nil {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	f.Seek(0, io.SeekStart)
	var isImage bool
	if _, err = xfile.ImageType(f); err == nil {
		isImage = true
	}
	dir := fmt.Sprintf("./assets/files/%d%d/%d/", time.Now().Year(), time.Now().Month(), time.Now().Day())
	dst := dir + filename + "." + xfile.Ext(file.Filename)
	if xfile.Exist(dst) {
		app.Success(c, app.Conf.Url+dst[1:])
		return
	}
	thumbDst := dir + filename + "_small." + xfile.Ext(file.Filename)
	// 上传文件至指定的完整文件路径
	if err = c.SaveUploadedFile(file, dst); err != nil {
		logger.Warnf("[api.upload] save file err: %v", err)
		app.Error(c, errno.ErrInternalServer)
		return
	}
	//检测图片类型，然后压缩
	if isImage {
		if err = thumb(dst, thumbDst); err != nil {
			logger.Warnf("[api.upload] thumb image err: %v", err)
			app.Error(c, errno.ErrInternalServer)
			return
		}
	}

	app.Success(c, app.Conf.Url+dst[1:])
}

// thumb 生成缩略图
func thumb(filepath, thumbPath string) error {
	if xfile.Exist(thumbPath) {
		return nil
	}
	src, err := imaging.Open(filepath)
	if err != nil {
		return err
	}
	dstImageFill := imaging.Fit(src, 200, 200, imaging.Lanczos)
	return imaging.Save(dstImageFill, thumbPath)
}
