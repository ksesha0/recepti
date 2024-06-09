package recipe

import (
	"fmt"
)

type Recipe struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

func DisplayRecipe(r Recipe) {
	fmt.Printf("Рецепт: %s\n", r.Name)
	totalWeight := 0.0
	totalCalories := 0.0
	totalProteins := 0.0
	totalFats := 0.0
	totalCarbohydrates := 0.0

	for _, ingredient := range r.Ingredients {
		fmt.Printf("Ингредиент: %s, Вес: %.2f г, Калории: %.2f, Белки: %.2f г, Жиры: %.2f г, Углеводы: %.2f г\n",
			ingredient.Name, ingredient.Weight, ingredient.Calories, ingredient.Proteins, ingredient.Fats, ingredient.Carbohydrates)
		totalWeight += ingredient.Weight
		totalCalories += ingredient.Calories
		totalProteins += ingredient.Proteins
		totalFats += ingredient.Fats
		totalCarbohydrates += ingredient.Carbohydrates
	}

	fmt.Printf("КБЖУ блюда '%s':\n", r.Name)
	fmt.Printf("Вес блюда: %.2f г\n", totalWeight)
	fmt.Printf("Калории: %.2f\n", totalCalories)
	fmt.Printf("Белки: %.2f г\n", totalProteins)
	fmt.Printf("Жиры: %.2f г\n", totalFats)
	fmt.Printf("Углеводы: %.2f г\n", totalCarbohydrates)
	if totalWeight > 0 {
		fmt.Printf("КБЖУ на 100 грамм: %.2f ккал, %.2f г белков, %.2f г жиров, %.2f г углеводов\n",
			totalCalories/totalWeight*100, totalProteins/totalWeight*100, totalFats/totalWeight*100, totalCarbohydrates/totalWeight*100)
	}

	fmt.Println("Сколько грамм блюда вы хотите приготовить?")
	var desiredWeight float64
	fmt.Scanln(&desiredWeight)

	if totalWeight > 0 && desiredWeight > 0 {
		factor := desiredWeight / totalWeight
		fmt.Printf("КБЖУ для %.2f грамм блюда '%s':\n", desiredWeight, r.Name)
		fmt.Printf("Калории: %.2f\n", totalCalories*factor)
		fmt.Printf("Белки: %.2f г\n", totalProteins*factor)
		fmt.Printf("Жиры: %.2f г\n", totalFats*factor)
		fmt.Printf("Углеводы: %.2f г\n", totalCarbohydrates*factor)
	} else {
		fmt.Println("Некорректный вес блюда.")
	}
}
