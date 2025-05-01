package test

import (
	"path/filepath"
	"testing"
	"tubes2/search"
	"tubes2/utils"
)

func TestBFS(t *testing.T) {
	jsonPath := filepath.Join("test", "fixture.json")
	g, err := utils.LoadAndBuildGraph(jsonPath)
	if err != nil {
		t.Fatalf("LoadAndBuildGraph failed: %v", err)
	}

	path, visited, _ := search.BFS(g, "A", "C")
	if len(path) == 0 {
		t.Fatal("Expected a non-empty path for BFS")
	}
	if path[len(path)-1] != "C" {
		t.Errorf("BFS target mismatch: got %v", path)
	}
	if visited <= 0 {
		t.Error("BFS visited should be > 0")
	}
}

func TestDFS(t *testing.T) {
	jsonPath := filepath.Join("test", "fixture.json")
	g, err := utils.LoadAndBuildGraph(jsonPath)
	if err != nil {
		t.Fatalf("LoadAndBuildGraph failed: %v", err)
	}

	path, visited, _ := search.DFS(g, "A", "C")
	if len(path) == 0 {
		t.Fatal("Expected a non-empty path for DFS")
	}
	if path[len(path)-1] != "C" {
		t.Errorf("DFS target mismatch: got %v", path)
	}
	if visited <= 0 {
		t.Error("DFS visited should be > 0")
	}
}
