const StatsPanel = ({ data }) => {
  if (!data || data.length === 0 || !data[0].recipe) {
    return <div>No recipe data available</div>; 
  }
  const result = data[0];
  return (
    <div>
      <h2>Stats Panel</h2>
      <p>Visited Nodes: {result.nodesVisited}</p>
      <p>Time Taken: {result.duration}</p>
      <p>Number of Recipes: {result.recipe ? 1 : 0}</p>  
    </div>
  );
};

export default StatsPanel;