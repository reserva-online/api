package user_test

import (
	"context"
	"testing"

	"github.com/schedule-api/pkg/authentication"
	"github.com/schedule-api/pkg/docker"
	"github.com/schedule-api/pkg/server"
	"github.com/schedule-api/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestService_Login(t *testing.T) {
	pgDb := docker.NewPostgres().WithTestPort(t)
	pgDb.Start(t)
	defer pgDb.Stop()

	db, _ := server.NewTestDatabase(pgDb.GetPort())
	s := user.NewService(db)

	userID, _ := s.Save(context.Background(), user.User{Email: "test@test", Password: "123", Type: "USER"})
	tests := []struct {
		name    string
		want    string
		login   user.LoginUser
		wantErr bool
	}{
		{
			name:    "should login and generate a valid token",
			login:   user.LoginUser{Email: "test@test", Password: "123"},
			wantErr: false,
		},
		{
			name:    "should fail at login",
			login:   user.LoginUser{Email: "wrongtest@test", Password: "123"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := s.Login(context.Background(), tt.login)
			if !tt.wantErr {
				assert.NotEqual(t, "", result.Token)
				assert.Nil(t, err)

				gotUserID, err := authentication.GetTokenUser(result.Token)

				assert.Nil(t, err)
				assert.Equal(t, userID, gotUserID)
			} else {
				assert.Equal(t, 0, result.ID)
				assert.NotNil(t, err)
			}
		})
	}
}
