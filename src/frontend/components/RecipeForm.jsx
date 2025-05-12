// src/components/RecipeForm.jsx
import React, { useState } from "react";
import { fetchRecipes } from "../utils/api";
import RecipeTree from "./RecipeTree";
import StatsPanel from "./StatsPanel";

const RecipeForm = ({ setResult }) => {
  const [start, setStart] = useState("Water");
  const [target, setTarget] = useState("");
  const [algo, setAlgo] = useState("bfs");  // Defaultnya BFS
  const [mode, setMode] = useState("single");
  const [max, setMax] = useState(3);
  const [error, setError] = useState(null);
  const [recipeData, setRecipeDataState] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!target) {
      setError("Target element is required.");
      return;
    }

    try {
      const response = await fetchRecipes(start, target, algo, mode, max);
      setRecipeDataState(response);  // Set recipe data ke state lokal
      setResult(response);  // Set hasil ke parent (index.jsx)
      setError(null);  
    } catch (err) {
      setError("An error occurred while fetching the recipe.");
    }
  };

  const handleToggleChange = () => {
    setMode(prevMode => (prevMode === "multiple" ? "single" : "multiple"));
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
            {/* <option value="bidirectional">Bidirectional</option> */}
          </select>
        </label>
        <br />
        <div className="toggle-container">
          <label>Multiple Recipe:</label>
          <label className="switch">
            <input
              type="checkbox"  // Tetap menggunakan checkbox
              checked={mode === "multiple"}
              onChange={handleToggleChange}
            />
            <span className="slider"></span>  {/* Slider dari global.css */}
          </label>
        </div>
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

      {recipeData && (
        <>
          <RecipeTree data={recipeData} />
          <StatsPanel data={recipeData} />
        </>
      )}
    </div>
  );
};

export default RecipeForm;