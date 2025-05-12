// src/frontend/pages/index.jsx
import React, { useState } from "react";
import RecipeForm from "../components/RecipeForm";

export default function Home() {
  const [result, setResult] = useState(null);

  return (
    <div className="p-8">
      <h1 className="text-3xl font-bold mb-4">Little Alchemy 2 Recipe Finder</h1>
      {/* Panggil RecipeForm untuk menampilkan form pencarian */}
      <RecipeForm setResult={setResult} />
    </div>
  );
}