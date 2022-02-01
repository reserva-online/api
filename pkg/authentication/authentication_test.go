package authentication_test

import (
	"testing"

	"github.com/schedule-api/pkg/authentication"
)

const fakeToken = "faketoken-.-faketoken"

func TestHandleUserCreateToken(t *testing.T) {
	userId := 1
	token, err := authentication.CreateToken(userId)

	if err != nil {
		t.Errorf("create token got = '%v', want no errors", err)
	}
	if len(token) < 100 {
		t.Errorf("create token got = '%v', want token", token)
	}
}

func TestHandleUserAuthentication(t *testing.T) {
	token, _ := authentication.CreateToken(1)

	tests := []struct {
		name        string
		tokenString string
		wantErr     bool
		wantUserID  int
	}{
		{
			name:        "should be able to parse token and return userID 1",
			tokenString: token,
			wantUserID:  1,
		},
		{
			name:        "shouldnt be able to parse token",
			tokenString: fakeToken,
			wantUserID:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			userID, err := authentication.GetTokenUser(token)

			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error while getting token user: %v", err)
			}
			if tt.wantUserID != userID {
				t.Errorf("Unexpected userID want = %v got %v", tt.wantUserID, userID)
			}
		})
	}
}

func TestCreateRecoverToken(t *testing.T) {
	email := "teste@gmail.com"
	token, err := authentication.CreateRecoverToken(email)

	if err != nil {
		t.Errorf("create token got = '%v', want no errors", err)
	}
	if len(token) < 100 {
		t.Errorf("create token got = '%v', want token", token)
	}
}

func TestGetEmailRecoverToken(t *testing.T) {
	email := "teste@gmail.com"
	token, _ := authentication.CreateRecoverToken(email)

	tests := []struct {
		name        string
		tokenString string
		wantErr     bool
		wantEmail   string
	}{
		{
			name:        "should be able to parse token and return email",
			tokenString: token,
			wantEmail:   email,
			wantErr:     false,
		},
		{
			name:        "shouldnt be able to parse token",
			tokenString: fakeToken,
			wantEmail:   email,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			email, err := authentication.GetEmailRecoverToken(token)

			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error while getting token user: %v", err)
			}
			if tt.wantEmail != email {
				t.Errorf("Unexpected email want = %v got %v", tt.wantEmail, email)
			}
		})
	}
}
