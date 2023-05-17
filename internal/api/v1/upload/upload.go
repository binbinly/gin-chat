package upload

import (
	"fmt"
	"time"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/util/xfile"
	"github.com/binbinly/pkg/util/xhash"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/ecode"
	"gin-chat/pkg/app"
)

const ImageMaxSize = 4 << 20 //最大上传4MB

// Image 上传图片
// @Summary 上传图片
// @Description 上传图片
// @Tags 上传
// @Accept		json
// @Produce		json
// @Param Token header string true "用户令牌"
// @Param file formData file true "图片"
// @success 0 {object} app.Response{data=string} "调用成功结构"
// @Router /image [post]
func Image(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	//图片限制大小
	if file.Size > ImageMaxSize {
		app.Error(c, ecode.ErrUploadImageLimit)
	}
	//打开文件
	f, err := file.Open()
	if err != nil {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	//检测图片类型
	if _, err = xfile.ImageType(f); err != nil {
		app.Error(c, ecode.ErrUploadNotImage)
		return
	}
	//文件名取md5值
	filename, err := xhash.MD5Reader(f)
	if err != nil {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	dir := fmt.Sprintf("./assets/images/%d%d/%d/", time.Now().Year(), time.Now().Month(), time.Now().Day())
	dst := dir + filename + xfile.Ext(file.Filename)
	if xfile.Exist(dst) {
		app.Success(c, dst[1:])
		return
	}
	thumbDst := dir + filename + "_small" + xfile.Ext(file.Filename)
	// 上传文件至指定的完整文件路径
	if err = c.SaveUploadedFile(file, dst); err != nil {
		app.Error(c, errno.ErrInternalServer)
		return
	}
	if err = thumb(dst, thumbDst); err != nil {
		app.Error(c, errno.ErrInternalServer)
		return
	}
	app.Success(c, dst[1:])
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
