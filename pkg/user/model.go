package user

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Type     string `db:"type"`
}
