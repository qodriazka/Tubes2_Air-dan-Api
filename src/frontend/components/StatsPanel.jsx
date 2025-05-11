// src/components/StatsPanel.jsx
import React from "react";

const StatsPanel = ({ data }) => {
  return (
    <div>
      <h2>Stats Panel</h2>
      <p>Visited Nodes: {data.steps.length}</p>
      <p>Time Taken: {data.durations[0]}</p>
      <p>Number of Recipes: {data.recipes.length}</p>
    </div>
  );
};

export default StatsPanel;