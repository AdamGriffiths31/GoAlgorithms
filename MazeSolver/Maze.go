package mazesolver

type point struct {
	x, y int
}

var direction = [4]point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall string, current, end point, seen *[][]bool, path *[]point) bool {
	return false
}

func solve(maze []string, wall string, current, end point) *[]point {
	return nil
}
