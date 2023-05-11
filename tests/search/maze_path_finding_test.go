package search_test

import (
	"testing"

	ds "github.com/kalogs-c/all-go-rythm/data_structures"
	"github.com/kalogs-c/all-go-rythm/search"
)

func TestMazePathFinding(t *testing.T) {
	maze := []string{
		"#### ##",
		"#      ",
		"# ## ##",
	}

	path, err := search.MazePathFinding(maze, "#", ds.Vector2{X: 1, Y: 2}, ds.Vector2{X: 4, Y: 0})
	if err != nil {
		t.Error(err)
	}

	expectedPath := []ds.Vector2{
		{X: 1, Y: 2},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 4, Y: 1},
		{X: 4, Y: 0},
	}

	for i := range path {
		if path[i] != expectedPath[i] {
			t.Errorf("error finding maze path, expected %v but was %v\n", expectedPath, path)
		}
	}
}

func TestMazeWithoutSolution(t *testing.T) {
	maze := []string{
		"#### ##",
		"#  #   ",
		"# ## ##",
	}

	_, err := search.MazePathFinding(maze, "#", ds.Vector2{X: 1, Y: 2}, ds.Vector2{X: 4, Y: 0})
	if err.Error() != "No solution" {
		t.Error(err)
	}
}
