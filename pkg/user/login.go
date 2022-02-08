package user

import (
	"context"
	"fmt"

	"github.com/schedule-api/pkg/authentication"
)

type LoginResponse struct {
	ID    int
	Token string
	Type  string
}

func (s *Service) Login(ctx context.Context, loginUser LoginUser) (LoginResponse, error) {
	var userItem User
	err := s.db.GetContext(ctx, &userItem, "SELECT * FROM users WHERE email = $1", loginUser.Email)
	if err != nil || loginUser.Password != userItem.Password {
		return LoginResponse{}, fmt.Errorf("credentials invalid")
	}

	token, err := authentication.CreateToken(userItem.ID)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("credentials invalid")
	}
	response := LoginResponse{
		ID:    userItem.ID,
		Type:  userItem.Type,
		Token: token,
	}

	return response, nil
}
