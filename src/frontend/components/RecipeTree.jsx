// src/components/RecipeTree.jsx
import React from "react";

// Fungsi untuk merender tree secara rekursif dengan kotak dan garis
const renderTree = (node, level = 0) => {
  if (!node) return null;

  const children = node.combines || [];

  return (
    <div className="recipe-tree-container" style={{ marginLeft: level * 50 }}>
      {/* Menampilkan kotak dengan nama elemen */}
      <div className="recipe-box">{node.name}</div>
      
      {/* Render children jika ada */}
      {children.length > 0 && (
        <div className="recipe-tree-branch">
          {/* Menampilkan garis vertikal antara elemen dan anak-anaknya */}
          <div className="line" style={{ height: `${children.length * 30}px` }}></div>
          
          {/* Menampilkan anak-anak dengan panah penghubung */}
          {children.map((combine, index) => (
            <div key={index} style={{ position: "relative" }}>
              {/* Garis penghubung ke bawah */}
              <div className="line" style={{ height: "20px" }}></div>
              {/* Panah mengarah ke bawah */}
              <div className="arrow" style={{ top: "20px", left: "35px" }}></div>
              {renderTree(combine, level + 1)}
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

const RecipeTree = ({ data }) => {
  if (!data || data.length === 0) {
    return <div>No data available</div>;  // Jika data kosong
  }

  const rootNode = data[0].recipe;  // Ambil recipe pertama dari data yang diterima

  return (
    <div>
      <h3>Recipe Tree</h3>
      <div>{renderTree(rootNode)}</div>  {/* Render pohon resep */}
    </div>
  );
};

export default RecipeTree;