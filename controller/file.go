package controller

import (
	"log"
	"time"

	"toilacoder/constant"
	"toilacoder/helper"
	"toilacoder/model"

	"github.com/kataras/iris"
)

func (c *Controller) UploadImage(ctx iris.Context) {
	auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	if !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	var uploadImage model.UploadImage
	err := ctx.ReadJSON(&uploadImage)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Định dạng sai")
		return
	}

	var image model.Image
	image.Title = uploadImage.Title
	image.Source = uploadImage.Url
	image.Filename = uploadImage.Filename
	image.CreatedAt = time.Now()

	if uploadImage.Replace == 0 {
		err = c.DB.Insert(&image)
		if err != nil {
			log.Println("Tên file đã bị trùng:", err)
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON("Tên file đã bị trùng")
			return
		}
	}

	if err := helper.DownloadFile(c.Config.ImageFolder+uploadImage.Filename, uploadImage.Url); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Upload ảnh lỗi")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(image.Filename)
}

func (c *Controller) DeleteImage(ctx iris.Context) {
	var deleteImage model.DeleteImage
	err := ctx.ReadJSON(&deleteImage)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Định dạng sai")
		return
	}

	var image model.Image
	image.Id = deleteImage.Id
	_, err = c.DB.Model(&image).WherePK().Delete()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Xoa image lỗi")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Xoá image thành công")
}
