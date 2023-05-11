package search

import (
	"errors"

	ds "github.com/kalogs-c/all-go-rythm/data_structures"
)

var POSSIBLE_DIR = []ds.Vector2{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}

func walk(
	maze []string,
	wall string,
	current ds.Vector2,
	goal ds.Vector2,
	visited map[ds.Vector2]bool,
	path *ds.Stack[ds.Vector2],
) bool {
	// Out of boundry
	if current.X < 0 || current.X >= len(maze[0]) || current.Y < 0 || current.Y >= len(maze) {
		return false
	}

	// Hitted a wall
	if string(maze[current.Y][current.X]) == wall {
		return false
	}

	// Reached the goal
	if current.X == goal.X && current.Y == goal.Y {
		return true
	}

	if visited[current] {
		return false
	}

	path.Push(current)

	for _, dir := range POSSIBLE_DIR {
		visited[current] = true
		next := ds.Vector2{
			X: current.X + dir.X,
			Y: current.Y + dir.Y,
		}
		if walk(maze, wall, next, goal, visited, path) {
			return true
		}
	}

	path.Pop()
	// Does not have a solution
	return false
}

func MazePathFinding(maze []string, wall string, start ds.Vector2, goal ds.Vector2) ([]ds.Vector2, error) {
	visited := make(map[ds.Vector2]bool)
	path := ds.NewStack[ds.Vector2]()

	walk(maze, wall, start, goal, visited, path)
	if path.Length() == 0 {
		return nil, errors.New("No solution")
	}

	pathSlice := path.ToSlice()

	return pathSlice, nil
}
