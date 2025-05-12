import React, { useEffect, useRef } from "react";
import mermaid from "mermaid";

const RecipeTree = ({ data }) => {
  const mermaidRef = useRef(null);

  useEffect(() => {
    if (data && data.length > 0) {
      if (mermaidRef.current) {
        mermaidRef.current.innerHTML = ''; 
      }
      data.forEach((recipeData, index) => {
        const rootNode = recipeData.recipe;
        let mermaidGraph = "graph TB; "; 
        const generateMermaidGraph = (node, parentId = null, index = 0) => {
          if (!node) return "";
          const nodeId = `${node.name.replace(/\s+/g, '').toLowerCase()}_${parentId ? parentId : 'root'}_${index}`;
          let graph = `${nodeId}[${node.name}]; `; 
          if (parentId) {
            graph += `${parentId} --> ${nodeId}; `;
          }
          if (node.combines && node.combines.length > 0) {
            node.combines.forEach((combine, idx) => {
              graph += generateMermaidGraph(combine, nodeId, idx);  
            });
          }
          return graph;
        };

        mermaidGraph += generateMermaidGraph(rootNode);
        const recipeDiv = document.createElement('div');
        recipeDiv.classList.add('mermaid'); 
        recipeDiv.innerHTML = mermaidGraph;  
        mermaidRef.current.appendChild(recipeDiv);
      });
      mermaid.initialize({ startOnLoad: true });
      mermaid.contentLoaded();
    }
  }, [data]);
  if (!data || data.length === 0) {
    return <div>No recipe data available</div>; 
  }
  return (
    <div>
      <h3>Recipe Tree</h3>
      <div ref={mermaidRef}></div> 
    </div>
  );
};

export default RecipeTree;