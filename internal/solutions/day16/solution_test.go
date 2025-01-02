package day16_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day16"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped long test because -short was set")
	}

	t.Parallel()

	s, err := day16.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		maze       [][]day16.Tile
		start, end map2d.Position

		// outputs
		expectedPaths int
		expectedScore int
		expectedSeats int
	}{
		"day's 1st example": {
			maze: [][]day16.Tile{
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.', '.', 'E', '#'},
				{'#', '.', '#', '.', '#', '#', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '.', '.', '.', '.', '.', '#', '.', '#', '.', '.', '.', '#', '.', '#'},
				{'#', '.', '#', '#', '#', '.', '#', '#', '#', '#', '#', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '.', '.', '.', '.', '.', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '#', '#', '#', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '#'},
				{'#', '#', '#', '.', '#', '.', '#', '#', '#', '#', '#', '.', '#', '.', '#'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.', '#', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '#', '#', '.', '#', '.', '#', '.', '#'},
				{'#', '.', '.', '.', '.', '.', '#', '.', '.', '.', '#', '.', '#', '.', '#'},
				{'#', '.', '#', '#', '#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#'},
				{'#', 'S', '.', '.', '#', '.', '.', '.', '.', '.', '#', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
			},
			start:         map2d.Position{X: 1, Y: 13},
			end:           map2d.Position{X: 13, Y: 1},
			expectedPaths: 3,
			expectedScore: 7036,
			expectedSeats: 45,
		},
		"day's 2nd example": {
			maze: [][]day16.Tile{
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.', 'E', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '#', '#', '#', '.', '#', '.', '#', '.', '#'},
				{'#', '.', '.', '.', '#', '.', '#', '.', '#', '.', '.', '.', '.', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '#', '#', '#', '#', '.', '#'},
				{'#', '.', '#', '.', '.', '.', '#', '.', '#', '.', '#', '.', '.', '.', '.', '.', '#'},
				{'#', '.', '#', '.', '#', '#', '#', '#', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.', '#'},
				{'#', '.', '#', '.', '#', '#', '#', '.', '#', '#', '#', '#', '#', '.', '#', '#', '#'},
				{'#', '.', '#', '.', '#', '.', '.', '.', '#', '.', '.', '.', '.', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '#', '#', '#', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '#', '#', '#', '#', '#', '#', '#', '#', '.', '#'},
				{'#', 'S', '#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
			},
			start:         map2d.Position{X: 1, Y: 15},
			end:           map2d.Position{X: 15, Y: 1},
			expectedPaths: 2,
			expectedScore: 11048,
			expectedSeats: 64,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day16.Solve(tt.maze, tt.start, tt.end)
			assert.Equal(t, tt.expectedPaths, len(got))
			for _, path := range got {
				assert.Equal(t, tt.expectedScore, path.Score)
			}

			gotSeats := day16.CountSeats(got)
			assert.Equal(t, tt.expectedSeats, gotSeats)
		})
	}
}
