package repositories

import (
	"database/sql"
	"test-1/internal/database"
	"test-1/internal/entities"
)

type UserRepository struct {
	database.BaseSQLRepository[entities.User]
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		BaseSQLRepository: database.BaseSQLRepository[entities.User]{DB: db},
	}
}

func (repo *UserRepository) GetAllUsers() ([]*entities.User, error) {
	query := "SELECT id, username, email FROM users"
	mapRow := func(rows *sql.Rows, user *entities.User) error {
		return rows.Scan(&user.ID, &user.Username, &user.Email)
	}
	return repo.SelectMultiple(mapRow, query)
}

func (repo *UserRepository) GetUserByID(id int) (*entities.User, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	mapRow := func(row *sql.Row, user *entities.User) error {
		return row.Scan(&user.ID, &user.Username, &user.Email)
	}
	return repo.SelectSingle(mapRow, query, id)
}

func (repo *UserRepository) CreateUser(user *entities.User) error {
	query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)"
	_, err := repo.Insert(query, user.Username, user.Email, user.Password)
	return err
}

func (repo *UserRepository) UpdateUser(user *entities.User) error {
	query := "UPDATE users SET username = $1, email = $2 WHERE id = $3"
	_, err := repo.ExecuteQuery(query, user.Username, user.Email, user.ID)
	return err
}
