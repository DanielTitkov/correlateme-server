package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) UserCount() (int, error) {
	return r.client.User.Query().Count(context.TODO())
}

func (r *EntgoRepository) GetUserByUsername(username string) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) GetUserByID(id int) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) CreateUser(u *domain.User) (*domain.User, error) {
	user, err := r.client.User.
		Create().
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPasswordHash(u.PasswordHash).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) GetUserCount() (int, error) {
	return r.client.User.Query().Count(context.Background())
}

func entToDomainUser(user *ent.User) *domain.User {
	return &domain.User{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}
