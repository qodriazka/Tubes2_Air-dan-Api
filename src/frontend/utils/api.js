export async function fetchRecipes(start, target, algo, mode, max) {
  try {
    const response = await fetch(
      `http://localhost:8080/search?start=${encodeURIComponent(start)}&target=${encodeURIComponent(target)}&algo=${encodeURIComponent(algo)}&mode=${encodeURIComponent(mode)}&max=${max}`
    );
    const data = await response.json();
    return data; // Mengembalikan seluruh data yang diterima dari backend
  } catch (error) {
    console.error("Error fetching recipes:", error);
    return { path: [], visited: 0, time_ms: 0 }; // Return default empty values if error
  }
}