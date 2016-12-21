https://leetcode.com/problems/island-perimeter/

Thoughts:

1) There is exactly one piece of 'land'
2) Iterate thru 2d array until you find that land, then expand in four directions recursively to explore all land
3) Hash table to save states you've visited so that you don't repeat / loop
4) How to calculate 'perimeter' of the island? It SEEMS like that's basically going to be 1 unit for every cell of water that a cell of land touches (including edges)
5) So every time a cell of land does not have a cell of land next to it (in a given direction), that's one unit of perimeter.

NOTES:

Arrays appear to be of fixed size
Slices are not

