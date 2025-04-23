export async function fetchRecipes(element) {
    try {
      const response = await fetch(`http://localhost:8080/search?element=${encodeURIComponent(element)}`);
      const data = await response.json();
      return data.recipes; // Pastikan backend balikin { recipes: [...] }
    } catch (error) {
      console.error("Error fetching recipes:", error);
      return [];
    }
  }
  