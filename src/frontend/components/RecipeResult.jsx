// src/components/RecipeResult.jsx
import React from "react";

const RecipeResult = ({ data }) => {
  return (
    <div>
      <h2>Recipe Results</h2>
      <p>Visited nodes: {data.steps.join(', ')}</p>  {/* Menampilkan langkah */}
      <p>Path: {data.recipes.map((recipe) => recipe.name).join(" → ")}</p>  {/* Menampilkan path resep */}
      <p>Time: {data.durations.join(', ')}</p>  {/* Menampilkan waktu */}
    </div>
  );
};

export default RecipeResult;