import React, { useState } from "react";
import { fetchRecipes } from "../utils/api";
import RecipeTree from "./RecipeTree";
import StatsPanel from "./StatsPanel";

const RecipeForm = ({ setResult }) => {
  const [target, setTarget] = useState("");
  const [algo, setAlgo] = useState("bfs");  // Defaultnya BFS
  const [mode, setMode] = useState("single");
  const [max, setMax] = useState(3);
  const [error, setError] = useState(null);
  const [recipeData, setRecipeDataState] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!target) {
      setError("Masukin dulu targetnya!");
      return;
    }

    try {
      const response = await fetchRecipes(target, algo, mode, max);  
      setRecipeDataState(response);  
      setResult(response); 
      setError(null);  
    } catch (err) {
      setError("An error occurred while fetching the recipe.");
    }
  };

  const handleToggleChange = () => {
    setMode(prevMode => (prevMode === "multiple" ? "single" : "multiple"));
  };

  // Function to handle algorithm button selection
  const handleAlgoClick = (algorithm) => {
    setAlgo(algorithm);
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          Target Element:
          <input
            type="text"
            value={target}
            onChange={(e) => setTarget(e.target.value)}
          />
        </label>
        <br />

        <label>Algoritma:</label>
        <div>
          <button
            type="button"
            className={`algo-btn ${algo === "bfs" ? "active" : ""}`}
            onClick={() => handleAlgoClick("bfs")}
          >
            BFS
          </button>
          <button
            type="button"
            className={`algo-btn ${algo === "dfs" ? "active" : ""}`}
            onClick={() => handleAlgoClick("dfs")}
          >
            DFS
          </button>
          <button
            type="button"
            className={`algo-btn ${algo === "bidirectional" ? "active" : ""}`}
            onClick={() => handleAlgoClick("bidirectional")}
          >
            Bidirectional
          </button>
        </div>
        <br />

        <div className="toggle-container">
          <label>Multiple Recipe:</label>
          <label className="switch">
            <input
              type="checkbox"
              checked={mode === "multiple"}
              onChange={handleToggleChange}
            />
            <span className="slider"></span>  
          </label>
        </div>
        <br />
        
        {mode === "multiple" && (
          <>
            <label>
              Max Recipes:
              <input
                type="range"
                min="1"
                max="6"
                value={max}
                onChange={(e) => setMax(e.target.value)}
                className="recipe-slider"
              />
              <span>{max}</span>
            </label>
            <br />
          </>
        )}
        
        <button className="btn">Search</button>
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
