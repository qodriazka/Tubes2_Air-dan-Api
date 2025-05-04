import React from "react";

const RecipeResult = ({ data }) => {
  return (
    <div>
      <h2>Recipe Results</h2>
      <p>Visited nodes: {data.visited}</p>
      <p>Path: {data.path.join(" → ")}</p>
      <p>Time: {data.time_ms} ms</p>
    </div>
  );
};

export default RecipeResult;