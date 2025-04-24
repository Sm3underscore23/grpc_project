package model

import "time"

type UserInfoPublic struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserInfoPrivate struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password string `db:"password"`
}

type UpdateUserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
}

type User struct {
	Id              int64 `db:"id"`
	UserInfoPrivate `db:""`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
