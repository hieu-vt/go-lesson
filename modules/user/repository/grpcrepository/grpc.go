package grpcrepository

import (
	"context"
	"lesson-5-goland/common"
	user "lesson-5-goland/proto/userproto"
)

type UserStore interface {
	GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}

type grpcStore struct {
	store UserStore
}

func NewGrpcUserRepository(store UserStore) *grpcStore {
	return &grpcStore{store: store}
}

func (s *grpcStore) GetUserByIds(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	userIds := make([]int, len(request.GetUserIds()))

	for i := range userIds {
		userIds[i] = int(request.GetUserIds()[i])
	}

	rs, err := s.store.GetUsers(ctx, userIds)

	if err != nil {
		return nil, err
	}

	users := make([]*user.User, len(rs))

	for i, item := range rs {
		users[i] = &user.User{
			Id:        int32(item.Id),
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Role:      item.Role,
		}
	}

	return &user.UserResponse{Users: users}, nil
}
