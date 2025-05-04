import React, { useState } from "react";
import { useRouter } from "next/router";
import { fetchRecipes } from "../utils/api"; 

const RecipeForm = ({ setRecipeData }) => {
  const [start, setStart] = useState("Water");
  const [target, setTarget] = useState("");
  const [algo, setAlgo] = useState("bfs");
  const [mode, setMode] = useState("multiple");
  const [max, setMax] = useState(3);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!target) {
      setError("Target element is required.");
      return;
    }

    try {
      const response = await fetchRecipes(start, target, algo, mode, max);
      setRecipeData(response); // Menyimpan data hasil pencarian ke state parent
      setError(null); 
    } catch (err) {
      setError("An error occurred while fetching the recipe.");
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          Start Element:
          <input
            type="text"
            value={start}
            onChange={(e) => setStart(e.target.value)}
          />
        </label>
        <br />
        <label>
          Target Element:
          <input
            type="text"
            value={target}
            onChange={(e) => setTarget(e.target.value)}
          />
        </label>
        <br />
        <label>
          Algorithm:
          <select value={algo} onChange={(e) => setAlgo(e.target.value)}>
            <option value="bfs">BFS</option>
            <option value="dfs">DFS</option>
          </select>
        </label>
        <br />
        <label>
          Mode:
          <select value={mode} onChange={(e) => setMode(e.target.value)}>
            <option value="multiple">Multiple</option>
            <option value="single">Single</option>
          </select>
        </label>
        <br />
        {mode === "multiple" && (
          <>
            <label>
              Max Recipes:
              <input
                type="number"
                value={max}
                onChange={(e) => setMax(parseInt(e.target.value))}
              />
            </label>
            <br />
          </>
        )}
        <button type="submit">Search</button>
      </form>

      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
};

export default RecipeForm;