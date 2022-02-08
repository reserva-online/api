package user

import (
	"context"
	"fmt"
)

func (s *Service) Save(ctx context.Context, data User) (int, error) {
	var validationUser User
	_ = s.db.GetContext(ctx, &validationUser, "SELECT * FROM users WHERE email = $1", data.Email)

	if validationUser.Email == data.Email {
		return 0, fmt.Errorf("User already exists")
	}

	var id int
	row, err := s.db.NamedQuery("INSERT INTO users(email, password, type) VALUES (:email, :password, :type) RETURNING id", data)
	if err != nil {
		return 0, err
	}
	if row.Next() {
		row.Scan(&id)
	}
	return id, nil

}
