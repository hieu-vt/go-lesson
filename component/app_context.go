package component

import (
	"gorm.io/gorm"
	"lesson-5-goland/component/uploadprovider/firebasestorage"
	"lesson-5-goland/component/uploadprovider/s3"
	"lesson-5-goland/pubsub"
	"lesson-5-goland/reddit"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() s3.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
	GetReddit() reddit.RedditEngine
	GetBuketFirebaseStorage() firebasestorage.UploadFirebaseStorageProvider
}

type appCtx struct {
	db             *gorm.DB
	provider       s3.UploadProvider
	secretKey      string
	pubsub         pubsub.Pubsub
	reddit         reddit.RedditEngine
	firebaseBucket firebasestorage.UploadFirebaseStorageProvider
}

func NewAppContext(
	db *gorm.DB,
	provider s3.UploadProvider,
	secretKey string,
	pubsub pubsub.Pubsub,
	reddit reddit.RedditEngine,
	firebaseBucket firebasestorage.UploadFirebaseStorageProvider,
) *appCtx {
	return &appCtx{
		db:             db,
		secretKey:      secretKey,
		provider:       provider,
		pubsub:         pubsub,
		reddit:         reddit,
		firebaseBucket: firebaseBucket,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() s3.UploadProvider {
	return ctx.provider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
	return ctx.pubsub
}

func (ctx *appCtx) GetReddit() reddit.RedditEngine {
	return ctx.reddit
}

func (ctx *appCtx) GetBuketFirebaseStorage() firebasestorage.UploadFirebaseStorageProvider {
	return ctx.firebaseBucket
}
