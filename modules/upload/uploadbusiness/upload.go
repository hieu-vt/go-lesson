package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"lesson-5-goland/common"
	"lesson-5-goland/component/uploadprovider/firebasestorage"
	"lesson-5-goland/component/uploadprovider/s3"
	"lesson-5-goland/modules/upload/uploadmodel"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type uploadBusiness struct {
	provider      s3.UploadProvider
	firebaseBuket firebasestorage.UploadFirebaseStorageProvider
}

func NewUploadBusiness(provider s3.UploadProvider, firebaseBuket firebasestorage.UploadFirebaseStorageProvider) *uploadBusiness {
	return &uploadBusiness{provider: provider, firebaseBuket: firebaseBuket}
}

func (biz *uploadBusiness) UploadFile(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s?alt=media", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	//img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	return img, nil
}

func (biz *uploadBusiness) UploadFileFirebase(ctx context.Context, file io.Reader, folder, fileName string) (*common.Image, error) {
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	img, err := biz.firebaseBuket.SaveFileUploaded(ctx, file, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = 0
	img.Height = 0
	//img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}

func getImageDimensionWithImaging(reader io.Reader) (int, int, error) {
	img, err := imaging.Decode(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Bounds().Dx(), img.Bounds().Dy(), nil
}
