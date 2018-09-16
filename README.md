Conway's (Golang) Game of Life
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

How to Play
-----

Install the dependencies with ```go get ./...```

Install the game globally with ```go install``` or locally with ```go build```. Consider renaming the executable by passing the ```-o``` flag

To view a game, run the executable with a duration, size and percentage of the cells to start the game alive (all integers)

```bash
gameoflife -s 20 -d 50 -p 55
```

for help check ```gameoflife -h```

To Do
---

- Build a better interface... or at least a prettier printout
- Properly test the random seeding
- Add initial setups which make for pretty pictures like on [http://www.conwaylife.com/](http://www.conwaylife.com/)

Notes to expand on
---

- Usually implementations of the game treat state in two steps. First, all the cells observe their neighbours and decide on their next state. Then that state is applied. Otherwise you'd get cells making their decision about what state to be in for period _x_ based on period _x_ states of cells earlier in the loop. I've intentionally not gone that route because I wanted to play around with iterating using channels and goroutines. That said, I'm fully aware that the result is arbitrary (there's no telling whether cells are making their observation based on the previous iteration or emerging results from the current one). I think it's an interesting point about this exercise. The rules don't say anything about whether time is supposed to be discrete or continuous. I might explore adding a trivial wait between the observation and the change of state and/or expanding the rules s.t. a neighbour has to have been in a given state for a certain amount of time when the observation is being made.
- My Structs are tightly coupled. Need to explore how Interfaces are used in go to work on that/ enter the generics in Go debate
- 