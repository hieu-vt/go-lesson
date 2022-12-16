package userbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component/tokenprovider"
	"lesson-5-goland/modules/user/usermodel"
)

type LoginStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type loginBiz struct {
	store         LoginStore
	hasher        Hasher
	tokenProvider tokenprovider.Provider
	expiry        int
}

func NewLoginBiz(store LoginStore, hasher Hasher, tokenProvider tokenprovider.Provider, expiry int) *loginBiz {
	return &loginBiz{store: store, hasher: hasher, tokenProvider: tokenProvider, expiry: expiry}
}

func (biz *loginBiz) Login(ctx context.Context, body *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": body.Email})

	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}

		return nil, err
	}

	if user.Id <= 0 {
		return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	checkedPassword := biz.hasher.Hash(body.Password + user.Salt)
	isMatch := checkedPassword == user.Password

	if !isMatch {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   string(user.Role),
	}

	token, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, err
	}

	return token, nil
}
