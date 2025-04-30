package test

import (
	"testing"
	"tubes2/search"
	"tubes2/utils"
)

func buildTestGraph() *utils.Graph {
	recipes := []utils.Recipe{
		{A: "Water", B: "Fire", Result: "Steam"},
		{A: "Earth", B: "Water", Result: "Mud"},
		{A: "Mud", B: "Fire", Result: "Brick"},
	}
	return utils.NewGraph(recipes)
}

func TestBFS(t *testing.T) {
	g := buildTestGraph()
	path, visited, _ := search.BFS(g, "Water", "Brick")
	expected := []string{"Water", "Mud", "Brick"}
	if len(path) != len(expected) {
		t.Fatalf("Expected path %v, got %v", expected, path)
	}
	for i := range path {
		if path[i] != expected[i] {
			t.Errorf("At index %d, expected %s, got %s", i, expected[i], path[i])
		}
	}
	if visited == 0 {
		t.Error("Visited count should be >0")
	}
}

func TestDFS(t *testing.T) {
	g := buildTestGraph()
	path, visited, _ := search.DFS(g, "Water", "Brick")
	if len(path) == 0 {
		t.Error("DFS did not find a path")
	}
	if visited == 0 {
		t.Error("Visited count should be >0")
	}
}
