package component

import (
	"gorm.io/gorm"
	"lesson-5-goland/component/uploadprovider"
	"lesson-5-goland/pubsub"
	"lesson-5-goland/reddit"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
	GetReddit() reddit.RedditEngine
}

type appCtx struct {
	db        *gorm.DB
	provider  uploadprovider.UploadProvider
	secretKey string
	pubsub    pubsub.Pubsub
	reddit    reddit.RedditEngine
}

func NewAppContext(
	db *gorm.DB,
	provider uploadprovider.UploadProvider,
	secretKey string,
	pubsub pubsub.Pubsub,
	reddit reddit.RedditEngine,
) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		provider:  provider,
		pubsub:    pubsub,
		reddit:    reddit,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
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
