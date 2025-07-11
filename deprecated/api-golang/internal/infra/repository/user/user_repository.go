package repository

import (
	"database/sql"

	"github.com/LuisMarchio03/nutri/internal/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Save(user *entity.User) error {
	stmt, err := r.Db.Prepare("INSERT INTO users (id, name, email, password, height, weight, age, gender, goal, total_calories, total_proteinas ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Height,
		user.Weight,
		user.Age,
		user.Gender,
		user.Goal,
		user.TotalCalories,
		user.TotalProteinas,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindAll() ([]*entity.User, error) {
	stmt, err := r.Db.Prepare("Select id, name, email, height, weight, age, gender, goal, total_calories, total_proteinas from users")
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
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Height,
			&user.Weight,
			&user.Age,
			&user.Gender,
			&user.Goal,
			&user.TotalCalories,
			&user.TotalProteinas,
		); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindInfosUser(userID string) (entity.User, error) {
	stmt, err := r.Db.Prepare("Select id, name, email, height, weight, age, gender, goal, total_calories, total_proteinas from users WHERE id = ?")
	if err != nil {
		return entity.User{}, err
	}
	row := stmt.QueryRow(userID)
	user := entity.User{}
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Height,
		&user.Weight,
		&user.Age,
		&user.Gender,
		&user.Goal,
		&user.TotalCalories,
		&user.TotalProteinas); err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(
	user *entity.User,
	userID string,
) error {
	stmt, err := r.Db.Prepare("UPDATE users SET name = ?, email = ?, password = ?, height = ?, weight = ?, age = ?, gender = ?, goal = ?, total_calories = ?, total_proteinas = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.Name,
		user.Email,
		user.Password,
		user.Height,
		user.Weight,
		user.Age,
		user.Gender,
		user.Goal,
		user.TotalCalories,
		user.TotalProteinas,
		userID,
	)
	if err != nil {
		return err
	}
	return nil

}
