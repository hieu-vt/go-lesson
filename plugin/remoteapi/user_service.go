package remoteapi

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/go-resty/resty/v2"
	"lesson-5-goland/common"
)

type userService struct {
	client     *resty.Client
	serviceURL string
	prefix     string
	log        logger.Logger
}

func NewUserApi(prefix string) *userService {
	return &userService{prefix: prefix}
}

func (u *userService) GetPrefix() string {
	return u.prefix
}

func (u *userService) Get() interface{} {
	return u
}

func (*userService) Name() string {
	return "user-api"
}

func (u *userService) InitFlags() {
	prefix := u.prefix
	if u.prefix != "" {
		prefix += "-"
	}

	flag.StringVar(&u.serviceURL, prefix+"secret", "http://localhost:8080", "Service url ex https://localhost:8081")
}

func (u *userService) Configure() error {
	u.client = resty.New()
	u.log = logger.GetCurrent().GetLogger(u.prefix)

	if u.serviceURL == "" {
		u.log.Errorln("Missing service URL")
		return errors.New("missing service URL")
	}

	return nil
}

func (userService) Run() error {
	return nil
}

func (userService) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (u *userService) GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error) {
	u.client = resty.New()

	type requestUserParam struct {
		Ids []int `json:"ids"`
	}

	type responseUser struct {
		Data []common.SimpleUser `json:"data"`
	}

	var result responseUser

	resp, err := u.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestUserParam{Ids: ids}).
		SetResult(&result).
		Post(fmt.Sprintf("%s/%s", u.serviceURL, "internal/get-users-by-ids"))

	if err != nil {
		//u.log.Infoln(err)
		return nil, err
	}

	if !resp.IsSuccess() {
		//u.log.Infoln(resp.RawResponse)
		return nil, errors.New(resp.RawResponse.Status)
	}

	for i := range result.Data {
		result.Data[i].GetRealId()
	}

	return result.Data, nil
}
