package repo

import (
	"context"

	"github.com/uptrace/bun"

	"usercenter/bc/user/model"
)

type UserRepo interface {
	Save(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	FindByPhone(ctx context.Context, phone string) (*model.User, error)
	FindByUID(ctx context.Context, uid string) (*model.User, error)
}

type userRepo struct {
	db *bun.DB
}

func (u *userRepo) Save(ctx context.Context, user *model.User) error {
	//TODO implement me
	u.db.NewUpdate()
	panic("implement me")
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) FindByPhone(ctx context.Context, phone string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) FindByUID(ctx context.Context, uid string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}
