// src/components/RecipeTree.jsx
import React from "react";

const renderTree = (node) => {
  if (!node) return null;
  if (!node.combines || node.combines.length === 0) {
    return <li>{node.name}</li>; // Jika node adalah elemen dasar (leaf)
  }
  return (
    <li>
      {node.name}
      <ul>
        {node.combines.map((combine, index) => (
          <li key={index}>
            <ul>
              <li>Left: {renderTree(combine.left)}</li>
              <li>Right: {renderTree(combine.right)}</li>
            </ul>
          </li>
        ))}
      </ul>
    </li>
  );
};

const RecipeTree = ({ data }) => {
  const rootNode = data.recipes[0]; // Ambil resep pertama (misalnya)
  return (
    <div>
      <h3>Recipe Tree</h3>
      <ul>{renderTree(rootNode)}</ul>
    </div>
  );
};

export default RecipeTree;