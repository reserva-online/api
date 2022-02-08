package docker_test

import (
	"testing"

	"github.com/schedule-api/pkg/docker"
	"github.com/schedule-api/pkg/server"
)

func TestPostgres_TestDatabaseCreation(t *testing.T) {
	pgDb := docker.NewPostgres().WithTestPort(t)
	pgDb.Start(t)
	defer pgDb.Stop()

	_, err := server.NewTestDatabase(pgDb.GetPort())

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "Should create a database for integration tests",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.NewTestDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
