Conway's Game of Go....errr... Life
===========

Description
-------

Conway's Game of Life in Go. Exercise in writing tested Go. And maybe concurrency... we'll see. Modelled after an exercise done at a Global Day of Code Retreat a few years ago. Commits are pedantic by design. 

Rules of the game
----

- Any live cell with fewer than two live neighbors dies, as if caused by underpopulation.
- Any live cell with more than three live neighbors dies, as if by overcrowding.
- Any live cell with two or three live neighbors lives on to the next generation.
- Any dead cell with exactly three live neighbors becomes a live cell.
