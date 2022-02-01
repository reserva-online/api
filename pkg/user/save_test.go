package user_test

import (
	"context"
	"testing"

	"github.com/schedule-api/pkg/docker"
	"github.com/schedule-api/pkg/server"
	"github.com/schedule-api/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestService_Save(t *testing.T) {
	pgDb := docker.NewPostgres().WithTestPort(t)
	pgDb.Start(t)
	defer pgDb.Stop()

	db, _ := server.NewTestDatabase(pgDb.GetPort())

	tests := []struct {
		name         string
		want         int
		toSave       user.User
		wantError    bool
		errorMessage string
	}{
		{
			name:      "should save",
			want:      1,
			toSave:    user.User{Email: "test@test", Password: "123", Type: "USER"},
			wantError: false,
		},
		{
			name:         "should not save because email already exists",
			toSave:       user.User{Email: "test@test", Password: "123", Type: "USER"},
			wantError:    true,
			errorMessage: "User already exists",
		},
		{
			name:      "should another user save",
			want:      2,
			toSave:    user.User{Email: "test2@test", Password: "123", Type: "USER"},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := user.NewService(db)
			got, err := s.Save(context.Background(), tt.toSave)

			if !tt.wantError {
				assert.Equal(t, tt.want, got)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.errorMessage, err.Error())
			}
		})
	}
}
