## Day 23

### Problem
https://adventofcode.com/2023/day/23

The problem just asks you to find the longest path in the given maze.

### Longest path in the maze

#### First input
* Finding the shortest path ih the maze is easy: just do BFS. But what about the longest path?
* There's no easy ways to find the longest path in a unweighted graph: https://en.wikipedia.org/wiki/Longest_path_problem
* By the way, seems it's almost one way? Are there any conditions making it easier to solve it?
* "never step onto the same tile twice"
* The maximum number of possible directions is 3 at most.
* Feeling like just brute force solution would work. Just list up all the possible paths.
* OK, I don't know what will come after this, but for now, let's do it.

v