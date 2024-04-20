package repository

import (
	"database/sql"

	"github.com/wizeline/CA-Microservices-Go/internal/entity"
)

type UserRepositoryPg struct {
	db *sql.DB
}

func NewUserRepositoryPg(db *sql.DB) UserRepositoryPg {
	return UserRepositoryPg{
		db: db,
	}
}

func (r UserRepositoryPg) Create(user entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (first_name, last_name, birthday, email, username, passwd) VALUES ($1, $2, $3, $4, $5, $6)",
		user.FirstName, user.LastName, user.BirthDay, user.Email, user.Username, user.Passwd,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepositoryPg) Read(id uint64) (entity.User, error) {
	var user entity.User
	row := r.db.QueryRow(`
		SELECT id, first_name, last_name, email, birthday,
			username, passwd, active, last_login,
			created_at, updated_at
		FROM users WHERE id = $1`, id)
	err := row.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.BirthDay,
		&user.Username, &user.Passwd, &user.Active, &user.LastLogin,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepositoryPg) ReadAll() ([]entity.User, error) {
	rows, err := r.db.Query(`
	SELECT id, first_name, last_name, email, birthday, 
		username, passwd, active, last_login,
		created_at, updated_at
	FROM users 
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]entity.User, 0)
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.BirthDay,
			&user.Username, &user.Passwd, &user.Active, &user.LastLogin,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r UserRepositoryPg) Update(user entity.User) error {
	_, err := r.db.Exec(`
		UPDATE users SET 
			first_name = $1,
			last_name = $2, 
			email = $3,
			birthday = $4,

			username = $5,
			passwd = $6, 
			active = $7,
			last_login = $8,
			updated_at = NOW()
		WHERE 
			id = $9`,
		user.FirstName, user.LastName, user.Email, user.BirthDay,
		user.Username, user.Passwd, user.Active, user.LastLogin,
		user.ID,
	)
	return err
}

func (r UserRepositoryPg) Delete(id uint64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
