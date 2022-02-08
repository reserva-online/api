package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/schedule-api/pkg/authentication"
)

type userInfo string

const userInfoKey userInfo = "userInfo"

type User struct {
	UserID int
}

func HandleAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			makeResponse(w, http.StatusUnauthorized, map[string]string{"message": "missing authorization token"})
			return
		}
		userID, err := authentication.GetTokenUser(token)
		if err != nil {
			makeResponse(w, http.StatusUnauthorized, map[string]string{"message": "invalid authorization token"})
			return
		}

		userIDContext := context.WithValue(r.Context(), userInfoKey, &User{UserID: userID})
		r = r.WithContext(userIDContext)
		h.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

/**
func getUserFromContext(ctx context.Context) *User {
	return ctx.Value(userInfoKey).(*User)
}

func getUserIdFromRequest(req *http.Request) int {
	user := getUserFromContext(req.Context())
	return user.UserID
}
*/
