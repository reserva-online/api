package user

import "context"

func (s *Service) GetById(ctx context.Context, id int) (User, error) {
	var userItem User
	err := s.db.Get(&userItem, "SELECT * FROM users WHERE id = $1", id)
	return userItem, err
}
