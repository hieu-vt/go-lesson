package component

import (
	"gorm.io/gorm"
	"lesson-5-goland/component/uploadprovider"
	"lesson-5-goland/pubsub"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db        *gorm.DB
	provider  uploadprovider.UploadProvider
	secretKey string
	pubsub    pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string, pubsub pubsub.Pubsub) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		provider:  provider,
		pubsub:    pubsub,
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
