package database

import (
	"database/sql"
	"github/LuisMarchio03/gointensivo/internal/accounts/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Save(user *entity.User) error {
	stmt, err := r.Db.Prepare("INSERT INTO users (id, name, age, cpf, birth, email, gender, password, street, city) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.ID,
		user.Name,
		user.Age,
		user.Cpf,
		user.Birth,
		user.Email,
		user.Gender,
		user.Password,
		user.Street,
		user.City,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindAll() ([]*entity.User, error) {
	stmt, err := r.Db.Prepare("Select id, name, age, cpf, birth, email, gender, street, city from users")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	users := []*entity.User{}
	for rows.Next() {
		user := new(entity.User)
		if err := rows.Scan(&user.ID,
			&user.Name,
			&user.Age,
			&user.Cpf,
			&user.Birth,
			&user.Email,
			&user.Gender,
			&user.Street,
			&user.City,
		); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users, nil
}
