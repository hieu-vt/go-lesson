package appcache

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"lesson-5-goland/modules/user/usermodel"
	"sync"
	"time"
)

const KeyUserAuthRedis = "user%d"

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type userAuthCache struct {
	realStore AuthenStore
	cache     AppCache
	once      *sync.Once
}

func NewUserAuthCache(realStore AuthenStore, cacheStore AppCache) *userAuthCache {
	return &userAuthCache{
		realStore: realStore,
		cache:     cacheStore,
		once:      new(sync.Once),
	}
}

func (uac *userAuthCache) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	var user usermodel.User
	userId := conditions["id"].(int)
	key := fmt.Sprintf(KeyUserAuthRedis, userId)
	//
	//_ = uac.cache.Get(ctx, key, &user)
	//
	//if user.Id > 0 {
	//	return &user, nil
	//} else {
	//	_ = uac.cache.Once(&cache.Item{
	//		Ctx:   ctx,
	//		Key:   key,
	//		Value: user,
	//		TTL:   time.Hour,
	//		Do: func(*cache.Item) (interface{}, error) {
	//			log.Println("get user from DB")
	//			u, err := uac.realStore.FindUser(ctx, conditions, moreInfo...)
	//
	//			if err != nil {
	//				return nil, err
	//			}
	//
	//			user = *u
	//
	//			return &u, nil
	//		},
	//	})
	//}
	//_ = uac.cache.Get(ctx, key, &user)
	//
	//return &user, nil
	_ = uac.cache.Get(ctx, key, &user)

	if user.Id > 0 {
		return &user, nil
	}

	if user.Id == 0 {
		var err error

		uac.once.Do(func() {
			log.Println("get user from DB")
			u, errDB := uac.realStore.FindUser(ctx, conditions)

			if err != nil {
				err = errDB
			} else {
				user = *u
				_ = uac.cache.Set(ctx, key, u, time.Hour*2)
			}
		})
	}

	_ = uac.cache.Get(ctx, key, &user)

	return &user, nil
}
