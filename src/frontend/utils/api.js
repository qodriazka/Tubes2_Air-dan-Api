export const fetchRecipes = async (target, algo, mode, max) => {
  const requestBody = {
    target: target,
    algorithm: algo,
    mode: mode,
    max_recipes: mode === "multiple" ? max : 1,
  };

  try {
    const response = await fetch("http://localhost:8080/search", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestBody),
    });

    if (!response.ok) {
      throw new Error("Failed to fetch recipes");
    }

    const data = await response.json();
    return data;  // Mengembalikan data yang diterima dari backend
  } catch (error) {
    console.error("Error fetching recipes:", error);
    throw error;
  }
};