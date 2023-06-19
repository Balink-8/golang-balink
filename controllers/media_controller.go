package controllers

import (
	"golang-img-api/dtos"
	"golang-img-api/models"
	"golang-img-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUpload() gin.HandlerFunc {
    return func(c *gin.Context) {
        //upload
        formFile, _, err := c.Request.FormFile("gambar")
        if err != nil {
            c.JSON(
                http.StatusInternalServerError,
                dtos.MediaDto{
                    StatusCode: http.StatusInternalServerError,
                    Message:    "error",
                    Data:       map[string]interface{}{"data": "Select a file to upload"},
                })
            return
        }

		gambarId := c.PostForm("gambar_id")
		produkId := c.PostForm("produk_id")

        uploadUrl, err := services.NewMediaUpload().FileUpload(models.GambarProduk{
			GambarId: gambarId,
			ProdukId: produkId,
			Gambar: formFile,
		})
        if err != nil {
            c.JSON(
                http.StatusInternalServerError,
                dtos.MediaDto{
                    StatusCode: http.StatusInternalServerError,
                    Message:    "error",
                    Data:       map[string]interface{}{"data": "Error uploading file"},
                })
            return
        }

        c.JSON(
            http.StatusOK,
            dtos.MediaDto{
                StatusCode: http.StatusOK,
                Message:    "success",
                Data:       map[string]interface{}{
					"gambar_id": gambarId,
					"produk_id": produkId,
					"gambar": uploadUrl,
				},
            })
    }
}
