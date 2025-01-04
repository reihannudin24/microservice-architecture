package repository

import (
	"Blast/internal/user/model"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) RegisterEmail(user model.User) (int64, error) {

	query := `insert into users_me (email) values ($1) RETURNING id`
	var id int64
	err := r.DB.QueryRow(query, user.Email).Scan(&id)

	return id, err
}

func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	query := "SELECT * FROM users_me WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.Id, &user.Email, &user.Contact, &user.Username, &user.Firstname, &user.Lastname, &user.Status, &user.CreatedAt, &user.UpdateAt)
	return &user, err
}
