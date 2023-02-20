package userbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/user/usermodel"
	"lesson-5-goland/plugin/jwtprovider"
	"lesson-5-goland/reddit"
)

type LoginStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type loginBiz struct {
	store         LoginStore
	hasher        Hasher
	tokenProvider jwtprovider.Provider
	expiry        int
	reddit        reddit.RedditEngine
}

func NewLoginBiz(store LoginStore, hasher Hasher, tokenProvider jwtprovider.Provider, expiry int, reddit reddit.RedditEngine) *loginBiz {
	return &loginBiz{store: store, hasher: hasher, tokenProvider: tokenProvider, expiry: expiry, reddit: reddit}
}

func (biz *loginBiz) Login(ctx context.Context, body *usermodel.UserLogin) (jwtprovider.Token, error) {
	ctxTrace, span := component.Tracer.Start(ctx, "user.biz.Login")
	defer span.End()
	user, err := biz.store.FindUser(ctxTrace, map[string]interface{}{"email": body.Email})

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

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: string(user.Role),
	}

	token, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, err
	}
	//biz.reddit.Save(fmt.Sprintf("%d", user.Id), user)

	return token, nil
}
