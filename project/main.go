package main

import (
	"fmt"
	"log"
	"project/database"
	"project/menu"
	"project/recipe"
)

func main() {
	fmt.Println("Добро пожаловать в программу 'Книга рецептов'!")

	err := database.CreateRecipeTable()
	if err != nil {
		log.Fatalf("Ошибка при создании таблиц: %v\n", err)
	}

	recipes, err := database.GetRecipes()
	if err != nil {
		fmt.Println("Ошибка загрузки рецептов:", err)
		recipes = []recipe.Recipe{}
	}

	recipes = menu.StartMenu(recipes)

	for _, r := range recipes {
		_, err := database.AddRecipe(r)
		if err != nil {
			fmt.Println("Ошибка сохранения рецепта:", err)
		}
	}
}
