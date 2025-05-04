/*
import React, { useState } from 'react'
import RecipeForm from '../components/RecipeForm'
import RecipeResult from '../components/RecipeResult'
import StatsPanel from '../components/StatsPanel'
*/

import React, { useState } from "react";
import { fetchRecipes } from "../utils/api";
import RecipeForm from "../components/RecipeForm";
import RecipeResult from "../components/RecipeResult";
import StatsPanel from "../components/StatsPanel";

export default function Home() {
  const [element, setElement] = useState("");
  const [recipes, setRecipes] = useState([]);
  const [result, setResult] = useState(null);

  const handleSearch = async () => {
    const result = await fetchRecipes(element);
    setRecipes(result);
  };

  return (
    <div className="p-8">
      <h1 className="text-3xl font-bold mb-4">Little Alchemy 2 Recipe Finder</h1>
      <RecipeForm setResult={setResult} />
      <input
        type="text"
        placeholder="Enter Element"
        value={element}
        onChange={(e) => setElement(e.target.value)}
        className="border p-2 mr-2"
      />
      <button onClick={handleSearch} className="bg-blue-500 text-white p-2">Search</button>
      {result && (
        <>
          <StatsPanel stats={result.stats} />
          <RecipeResult recipe={result.recipe} />
        </>
      )}
      <ul className="mt-4">
        {recipes.map((recipe, idx) => (
          <li key={idx}>{recipe}</li>
        ))}
      </ul>
    </div>
  );
}



/*
export default function IndexPage() {
  const [result, setResult] = useState(null)

  return (
    <main>
      <h1>Little Alchemy 2 Recipe Finder</h1>
      <RecipeForm setResult={setResult} />
      {result && <>
        <StatsPanel stats={result.stats} />
        <RecipeResult recipe={result.recipe} />
      </>}
    </main>
  )
}
*/
