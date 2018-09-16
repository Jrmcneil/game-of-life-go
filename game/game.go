package game

import "game-of-life-go/grid"

type Game struct {
    dimensions int
    grid *grid.Grid
}

func (game *Game) build(dimension int) *grid.Grid {
    game.grid = grid.NewGrid(dimension, dimension)
    game.grid.Fill()
    return game.grid
}