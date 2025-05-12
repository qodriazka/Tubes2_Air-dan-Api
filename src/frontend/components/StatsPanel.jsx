// src/components/StatsPanel.jsx
import React from "react";

const StatsPanel = ({ data }) => {
  // Memastikan data yang diterima memiliki struktur yang benar
  if (!data || data.length === 0 || !data[0].recipe) {
    return <div>No recipe data available</div>;  // Jika tidak ada data resep
  }

  // Ambil hasil pertama dari array data
  const result = data[0];

  return (
    <div>
      <h2>Stats Panel</h2>
      {/* Menampilkan nodesVisited */}
      <p>Visited Nodes: {result.nodesVisited}</p>
      {/* Menampilkan waktu dari result */}
      <p>Time Taken: {result.duration}</p>
      {/* Menampilkan jumlah resep */}
      <p>Number of Recipes: {result.recipe ? 1 : 0}</p>  {/* Hanya ada 1 resep dalam contoh */}
    </div>
  );
};

export default StatsPanel;