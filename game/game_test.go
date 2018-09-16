package game

import "testing"

func TestGameTakesParametersAndCreatesAGridOfCells(t *testing.T) {
    game := new(Game)
    game.build(5)
    grid := game.grid

    if  grid.GetCells() != 25 || grid.GetHeight() != 5 || grid.GetWidth() != 5 {
        t.Errorf("Initialised game should create a Grid of the correct size and number of Cells")
    }
}
