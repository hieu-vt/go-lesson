package ginupload

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/upload/uploadbusiness"
)

func UploadFile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		defer file.Close() // we can close here

		//dataBytes := make([]byte, fileHeader.Size)
		//if _, err := file.Read(dataBytes); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}

		//imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider(), appCtx.GetBuketFirebaseStorage())
		img, err := biz.UploadFileFirebase(c.Request.Context(), file, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
