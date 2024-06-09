CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredients (
    id SERIAL PRIMARY KEY,
    recipe_id INTEGER REFERENCES recipes(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    weight FLOAT NOT NULL,
    calories FLOAT NOT NULL,
    proteins FLOAT NOT NULL,
    fats FLOAT NOT NULL,
    carbohydrates FLOAT NOT NULL
);
