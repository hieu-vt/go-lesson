package firebasestorage

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"io"
	"lesson-5-goland/common"
	"strings"
)

type firebaseStorage struct {
	bucket                   *storage.BucketHandle
	ProjectFirebaseStorageID string
	StorageBucket            string
	FirebaseStorageDomain    string
}

func NewFirebaseStorage(
	ctx context.Context,
	ProjectFirebaseStorageID string,
	StorageBucket string,
	FirebaseStorageDomain string,
	CredentialsFileName string,
) *firebaseStorage {
	config := &firebase.Config{
		ProjectID:     ProjectFirebaseStorageID,
		StorageBucket: StorageBucket,
	}
	otp := option.WithCredentialsFile(CredentialsFileName)

	app, err := firebase.NewApp(ctx, config, otp)
	if err != nil {
		log.Infof("firebase.NewApp: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Infof("app.Storage: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Infof("client.DefaultBucket: %v", err)
	}

	return &firebaseStorage{
		bucket:                   bucket,
		ProjectFirebaseStorageID: ProjectFirebaseStorageID,
		StorageBucket:            StorageBucket,
		FirebaseStorageDomain:    FirebaseStorageDomain,
	}
}

func (b *firebaseStorage) SaveFileUploaded(ctx context.Context, file io.Reader, dst string) (*common.Image, error) {
	obj := b.bucket.Object(dst)
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Infof("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Infof("wc.Close: %v", err)
	}

	url, err := obj.Attrs(ctx)
	if err != nil {
		log.Infof("obj.Attrs: %v", err)
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s%s", b.FirebaseStorageDomain, strings.SplitAfterN(url.MediaLink, "v1", 3)[1]),
		CloudName: "firebase-storage",
	}

	return img, nil
}
