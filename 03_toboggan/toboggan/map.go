package toboggan

import "errors"

// A Map represents the local geology.
type Map struct {
	grid   []string
	Width  int
	Height int
}

// ParseMap returns a Map parsed from grid.
func ParseMap(grid []string) (*Map, error) {
	if len(grid) == 0 {
		return nil, errors.New("empty map")
	}

	width := len(grid[0])
	height := len(grid)

	for h := range grid {
		if len(grid[h]) != width {
			return nil, errors.New("invalid map")
		}
	}

	m := &Map{grid, width, height}

	return m, nil
}

func (m Map) location(w int, h int) rune {
	return rune(m.grid[h][w%m.Width])
}

// IsTree indicates if there is a tree on the Map at location w, h.
func (m Map) isTree(w int, h int) bool {
	l := m.location(w, h)
	return l == '#'
}

// TreesOnSlope returns the number of trees encountered on the map with the given slope.
func (m Map) TreesOnSlope(right int, down int) int {
	w, h := 0, 0
	trees := 0
	for h < m.Height-down {
		w += right
		h += down
		if m.isTree(w, h) {
			trees++
		}
	}

	return trees
}
