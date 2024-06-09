package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"project/recipe"
)

func LoadRecipes(filename string) ([]recipe.Recipe, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := ioutil.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var recipes []recipe.Recipe
	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func SaveRecipes(filename string, recipes []recipe.Recipe) error {
	data, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
