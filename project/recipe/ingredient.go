package recipe

type Ingredient struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Weight        float64 `json:"weight"`
	Calories      float64 `json:"calories"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}
