export const fetchRecipes = async (target, algo, mode, max) => {
  const requestBody = {
    target: target,
    algorithm: algo,
    mode: mode,
  };
  if (mode === "multiple") {
    requestBody.max_recipes = parseInt(max);
  }

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
    return data;  // Return data from the backend
  } catch (error) {
    console.error("Error fetching recipes:", error);
    throw error;
  }
};