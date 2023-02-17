package userbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/user/usermodel"
)

type RegisterUserStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	hasher Hasher
	store  RegisterUserStore
}

func NewRegisterBiz(store RegisterUserStore, hasher Hasher) *registerBiz {
	return &registerBiz{
		store:  store,
		hasher: hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserCreate, error) {
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	//if err != nil {
	//	if err != common.RecordNotFound {
	//		return nil, common.RecordNotFound
	//	}
	//	return nil, err
	//}

	if user != nil && user.Id > 0 {
		return nil, usermodel.ErrEmailExisted
	}

	data.Status = 1
	data.Salt = common.GenSalt(20)
	data.Password = biz.hasher.Hash(data.Password + data.Salt)

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	return data, nil
}
