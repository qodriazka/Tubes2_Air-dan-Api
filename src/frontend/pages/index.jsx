// src/frontend/pages/index.jsx
import React, { useState } from "react";
import { fetchRecipes } from "../utils/api";
import RecipeForm from "../components/RecipeForm";
import RecipeResult from "../components/RecipeResult";
import StatsPanel from "../components/StatsPanel";

export default function Home() {
  const [result, setResult] = useState(null);

  return (
    <div className="p-8">
      <h1 className="text-3xl font-bold mb-4">Little Alchemy 2 Recipe Finder</h1>
      <RecipeForm setResult={setResult} />
      
      {result && (
        <>
          <StatsPanel stats={result.stats} />
          <RecipeResult recipe={result.recipe} />
        </>
      )}
    </div>
  );
}