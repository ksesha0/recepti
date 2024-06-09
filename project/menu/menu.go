package menu

import (
	"fmt"
	"project/database"
	"project/recipe"
	"strconv"
)

func StartMenu(recipes []recipe.Recipe) []recipe.Recipe {
	for {
		fmt.Println("Что вы хотите сделать?")
		fmt.Println("1 - Посмотреть рецепты")
		fmt.Println("2 - Добавить рецепт")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displayRecipes(recipes)
		case 2:
			recipes = addRecipe(recipes)
			// Сохраняем рецепты в базу данных после добавления
			err := database.SaveRecipes("recipes.json", recipes)
			if err != nil {
				fmt.Println("Ошибка сохранения рецептов:", err)
			}
		default:
			fmt.Println("Некорректный выбор.")
		}

		fmt.Println("Хотите продолжить? (y/n)")
		var cont string
		fmt.Scanln(&cont)
		if cont != "y" && cont != "Y" {
			break
		}
	}
	return recipes
}

func displayRecipes(recipes []recipe.Recipe) {
	fmt.Println("Рецепты:")
	for i, r := range recipes {
		fmt.Printf("%d - %s\n", i+1, r.Name)
	}
	fmt.Println("Введите номер рецепта для подробного просмотра или 0 для возврата:")
	var recipeNumber int
	fmt.Scanln(&recipeNumber)
	if recipeNumber > 0 && recipeNumber <= len(recipes) {
		recipe.DisplayRecipe(recipes[recipeNumber-1])
	} else if recipeNumber != 0 {
		fmt.Println("Некорректный номер рецепта.")
	}
}

func addRecipe(recipes []recipe.Recipe) []recipe.Recipe {
	fmt.Println("Введите название рецепта:")
	var name string
	fmt.Scanln(&name)

	var ingredients []recipe.Ingredient
	numIngredients := getIntInput("Сколько ингредиентов у вас в рецепте?")

	for i := 0; i < numIngredients; i++ {
		fmt.Printf("Введите информацию об ингредиенте %d:\n", i+1)
		var ingredient recipe.Ingredient

		fmt.Println("Название ингредиента:")
		fmt.Scanln(&ingredient.Name)

		ingredient.Weight = getFloatInput("Вес (граммы):")
		ingredient.Calories = getFloatInput("Калории:")
		ingredient.Proteins = getFloatInput("Белки:")
		ingredient.Fats = getFloatInput("Жиры:")
		ingredient.Carbohydrates = getFloatInput("Углеводы:")

		ingredients = append(ingredients, ingredient)
	}

	r := recipe.Recipe{Name: name, Ingredients: ingredients}
	recipes = append(recipes, r)

	fmt.Println("Рецепт успешно добавлен!")
	recipe.DisplayRecipe(r)
	return recipes
}

func getFloatInput(prompt string) float64 {
	for {
		fmt.Println(prompt)
		var input string
		fmt.Scanln(&input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Некорректный ввод. Пожалуйста, введите число.")
			continue
		}
		return value
	}
}

func getIntInput(prompt string) int {
	for {
		fmt.Println(prompt)
		var input string
		fmt.Scanln(&input)
		value, err := strconv.Atoi(input)
		if err != nil || value < 0 {
			fmt.Println("Некорректный ввод. Пожалуйста, введите положительное целое число.")
			continue
		}
		return value
	}
}
