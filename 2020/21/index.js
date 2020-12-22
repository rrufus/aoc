const { readFileSync } = require("fs");
const input = readFileSync("input").toString();

const parsed = input
  .trim()
  .split("\n")
  .reduce(
    (memo, line) => {
      const [rawIngredientsStr, contains] = line.split("(contains ");
      const rawIngredients = rawIngredientsStr.trim().split(" ");
      const allergens = contains.split(")")[0].split(", ");

      const allergenToIngredients = allergens.reduce(
        (m, allergen) => ({
          ...m,
          [allergen]: m[allergen]
            ? rawIngredients.filter((ingredient) =>
                m[allergen].includes(ingredient)
              )
            : rawIngredients,
        }),
        memo.allergenToIngredients
      );

      const ingredientToFrequency = rawIngredients.reduce(
        (m, ingredient) => ({
          ...m,
          [ingredient]: m[ingredient] ? m[ingredient] + 1 : 1,
        }),
        memo.ingredientToFrequency
      );

      return { allergenToIngredients, ingredientToFrequency };
    },
    { allergenToIngredients: {}, ingredientToFrequency: {} }
  );

console.log(
  "Part 1: ",
  Object.keys(parsed.ingredientToFrequency)
    .filter((i) =>
      Object.values(parsed.allergenToIngredients).every(
        (ingredients) => !ingredients.includes(i)
      )
    )
    .reduce(
      (sum, ingredient) => (sum += parsed.ingredientToFrequency[ingredient]),
      0
    )
);

while (
  Object.values(parsed.allergenToIngredients).reduce(
    (sum, list) => (sum += list.length),
    0
  ) > Object.keys(parsed.allergenToIngredients).length
) {
  parsed.allergenToIngredients = Object.entries(
    parsed.allergenToIngredients
  ).reduce(
    (memo, [subjectAllergen, subjectIngredients]) =>
      subjectIngredients.length === 1
        ? Object.entries(memo).reduce(
            (memo, [allergen, ingredients]) => ({
              ...memo,
              [allergen]:
                allergen === subjectAllergen
                  ? ingredients
                  : ingredients.filter((i) => i !== subjectIngredients[0]),
            }),
            {}
          )
        : memo,
    parsed.allergenToIngredients
  );
}

console.log(
  "Part 2: ",
  Object.keys(parsed.allergenToIngredients)
    .sort()
    .map((a) => parsed.allergenToIngredients[a][0])
    .join(",")
);
