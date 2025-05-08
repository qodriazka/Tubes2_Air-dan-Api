package test

import (
	"path/filepath"
	"testing"
	"tubes2/search"
	"tubes2/utils"
)

func TestBFS(t *testing.T) {
	jsonPath := filepath.Join("test", "fixture.json")
	g, err := utils.NewGraphFromJSON(jsonPath)
	if err != nil {
		t.Fatalf("NewGraphFromJSON failed: %v", err)
	}

	path, visited, _ := search.BFS(g, "A")
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
	g, err := utils.NewGraphFromJSON(jsonPath)
	if err != nil {
		t.Fatalf("NewGraphFromJSON failed: %v", err)
	}

	path, visited, _ := search.DFS(g, "A")
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

func TestBidirectionalSearch(t *testing.T) {
	jsonPath := filepath.Join("test", "fixture.json")
	g, err := utils.NewGraphFromJSON(jsonPath)
	if err != nil {
		t.Fatalf("NewGraphFromJSON failed: %v", err)
	}

	path, visited, _ := search.BidirectionalSearch(g, "A")
	if len(path) == 0 {
		t.Fatal("Expected a non-empty path for Bidirectional Search")
	}
	if path[len(path)-1] != "C" {
		t.Errorf("Bidirectional Search target mismatch: got %v", path)
	}
	if visited <= 0 {
		t.Error("Bidirectional Search visited should be > 0")
	}
}
