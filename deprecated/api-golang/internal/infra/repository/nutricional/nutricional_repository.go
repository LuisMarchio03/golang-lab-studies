package repository

import (
	"database/sql"

	"github.com/LuisMarchio03/nutri/internal/entity"
)

type NutricionalRepository struct {
	Db *sql.DB
}

func NewNutricionalRepository(db *sql.DB) *NutricionalRepository {
	return &NutricionalRepository{Db: db}
}

func (r *NutricionalRepository) Save(n *entity.Nutricional) error {
	_, err := r.Db.Exec("INSERT INTO nutricional (id, id_taco, food_name, quantity, calories_unit, total_calories) VALUES (?, ?, ?, ?, ?, ?)", n.ID, n.IDTaco, n.FoodName, n.Quantity, n.CaloriesUnit, n.TotalCalories)
	if err != nil {
		return err
	}
	return nil
}

func (r *NutricionalRepository) FindAll() ([]*entity.Nutricional, error) {
	rows, err := r.Db.Query("SELECT id, id_taco, food_name, quantity, calories_unit, total_calories FROM nutricional")
	if err != nil {
		return nil, err
	}

	var nutricionals []*entity.Nutricional
	for rows.Next() {
		var nutricional entity.Nutricional
		err = rows.Scan(
			&nutricional.ID,
			&nutricional.IDTaco,
			&nutricional.FoodName,
			&nutricional.Quantity,
			&nutricional.CaloriesUnit,
			&nutricional.TotalCalories,
		)
		if err != nil {
			return nil, err
		}
		nutricionals = append(nutricionals, &nutricional)
	}
	return nutricionals, nil
}

func (r *NutricionalRepository) FindUnique(ID string) (*entity.Nutricional, error) {
	var nutricional entity.Nutricional
	err := r.Db.QueryRow("SELECT id, id_taco, food_name, quantity, calories_unit, total_calories FROM nutricional WHERE id = ?", ID).Scan(
		&nutricional.ID,
		&nutricional.IDTaco,
		&nutricional.FoodName,
		&nutricional.Quantity,
		&nutricional.CaloriesUnit,
		&nutricional.TotalCalories,
	)
	if err != nil {
		return nil, err
	}
	return &nutricional, nil
}

func (r *NutricionalRepository) Update(nutricional *entity.Nutricional) error {
	stmt, err := r.Db.Prepare("UPDATE nutricional SET id_taco = ?, food_name = ?, quantity = ?, calories_unit = ?, total_calories = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		nutricional.IDTaco,
		nutricional.FoodName,
		nutricional.Quantity,
		nutricional.CaloriesUnit,
		nutricional.TotalCalories,
		nutricional.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *NutricionalRepository) Delete(ID string) error {
	stmt, err := r.Db.Prepare("DELETE FROM nutricional WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}
	return nil
}
