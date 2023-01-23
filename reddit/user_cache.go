package reddit

import (
	"context"
	"fmt"
	"lesson-5-goland/modules/user/usermodel"
)

type UserStoreMySql interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type userCacheStore struct {
	store      RedditEngine
	storeMysql UserStoreMySql
}

func NewUserCache(store RedditEngine, storeMySql UserStoreMySql) *userCacheStore {
	return &userCacheStore{
		store:      store,
		storeMysql: storeMySql,
	}
}

func (uCache *userCacheStore) FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	id := condition["id"].(int)

	cacheUser := uCache.store.Get(fmt.Sprintf("%d", id))
	if cacheUser != nil {
		return cacheUser.(*usermodel.User), nil
	}
	// handle destroy left of reddit -- timer run after 1-3 h set nil
	return uCache.storeMysql.FindUser(ctx, condition)
}
