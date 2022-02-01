package user_test

import (
	"context"
	"testing"

	"github.com/schedule-api/pkg/docker"
	"github.com/schedule-api/pkg/server"
	"github.com/schedule-api/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestService_GetById(t *testing.T) {
	pgDb := docker.NewPostgres().WithTestPort(t)
	pgDb.Start(t)
	defer pgDb.Stop()

	db, _ := server.NewTestDatabase(pgDb.GetPort())

	s := user.NewService(db)
	savedUserId, _ := s.Save(context.Background(), user.User{ID: 1, Email: "test@test", Password: "123", Type: "USER"})

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "should get saved user",
			id:      savedUserId,
			wantErr: false,
		},
		{
			name:    "should not find user by id",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetById(context.Background(), tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.id, got.ID)
		})
	}
}
