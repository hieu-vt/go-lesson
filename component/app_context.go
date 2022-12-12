package component

import (
	"gorm.io/gorm"
	"lesson-5-goland/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db       *gorm.DB
	provider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		db:       db,
		provider: provider,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.provider
}
