package component

import (
	"gorm.io/gorm"
	"lesson-5-goland/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db        *gorm.DB
	provider  uploadprovider.UploadProvider
	secretKey string
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		provider:  provider,
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
