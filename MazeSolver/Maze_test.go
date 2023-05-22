package mazesolver

import "testing"

func TestSolve(t *testing.T) {
	mazeResult := []point{
		{10, 0},
		{10, 1},
		{10, 2},
		{10, 3},
		{10, 4},
		{9, 4},
		{8, 4},
		{7, 4},
		{6, 4},
		{5, 4},
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
		{1, 5},
	}

	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	wall := "x"
	current := point{10, 0}
	end := point{1, 5}

	path := solve(maze, wall, current, end)
	if len(*path) != len(mazeResult) {
		t.Errorf("Expected %d, got %d", len(mazeResult), len(*path))
	}
	for i := 0; i < len(*path); i++ {
		if (*path)[i] != mazeResult[i] {
			t.Errorf("Expected %v, got %v", mazeResult[i], (*path)[i])
		}
	}
}
