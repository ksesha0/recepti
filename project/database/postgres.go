package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"project/recipe"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка проверки соединения с базой данных: %v\n", err)
	}
}

func CreateRecipeTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS recipes (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS ingredients (
		id SERIAL PRIMARY KEY,
		recipe_id INTEGER REFERENCES recipes(id) ON DELETE CASCADE,
		name TEXT NOT NULL,
		weight FLOAT NOT NULL,
		calories FLOAT NOT NULL,
		proteins FLOAT NOT NULL,
		fats FLOAT NOT NULL,
		carbohydrates FLOAT NOT NULL
	);`

	_, err := db.Exec(query)
	return err
}

func AddRecipe(r recipe.Recipe) (int, error) {
	var recipeID int
	err := db.QueryRow("INSERT INTO recipes (name) VALUES ($1) RETURNING id", r.Name).Scan(&recipeID)
	if err != nil {
		return 0, err
	}

	for _, ingredient := range r.Ingredients {
		_, err = db.Exec("INSERT INTO ingredients (recipe_id, name, weight, calories, proteins, fats, carbohydrates) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			recipeID, ingredient.Name, ingredient.Weight, ingredient.Calories, ingredient.Proteins, ingredient.Fats, ingredient.Carbohydrates)
		if err != nil {
			return 0, err
		}
	}

	return recipeID, nil
}

func GetRecipes() ([]recipe.Recipe, error) {
	rows, err := db.Query("SELECT id, name FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []recipe.Recipe
	for rows.Next() {
		var r recipe.Recipe
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}

		ingredientRows, err := db.Query("SELECT id, name, weight, calories, proteins, fats, carbohydrates FROM ingredients WHERE recipe_id = $1", r.ID)
		if err != nil {
			return nil, err
		}
		defer ingredientRows.Close()

		for ingredientRows.Next() {
			var i recipe.Ingredient
			if err := ingredientRows.Scan(&i.ID, &i.Name, &i.Weight, &i.Calories, &i.Proteins, &i.Fats, &i.Carbohydrates); err != nil {
				return nil, err
			}
			r.Ingredients = append(r.Ingredients, i)
		}
		recipes = append(recipes, r)
	}

	return recipes, nil
}

func UpdateRecipe(r recipe.Recipe) error {
	_, err := db.Exec("UPDATE recipes SET name = $1 WHERE id = $2", r.Name, r.ID)
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM ingredients WHERE recipe_id = $1", r.ID)
	if err != nil {
		return err
	}

	for _, ingredient := range r.Ingredients {
		_, err = db.Exec("INSERT INTO ingredients (recipe_id, name, weight, calories, proteins, fats, carbohydrates) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			r.ID, ingredient.Name, ingredient.Weight, ingredient.Calories, ingredient.Proteins, ingredient.Fats, ingredient.Carbohydrates)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteRecipe(id int) error {
	_, err := db.Exec("DELETE FROM ingredients WHERE recipe_id = $1", id)
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM recipes WHERE id = $1", id)
	return err
}
