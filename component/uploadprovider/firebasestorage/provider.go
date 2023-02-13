package firebasestorage

import (
	"context"
	"io"
	"lesson-5-goland/common"
)

type UploadFirebaseStorageProvider interface {
	SaveFileUploaded(ctx context.Context, file io.Reader, dst string) (*common.Image, error)
}
