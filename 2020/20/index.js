const { readFileSync } = require("fs");
const input = readFileSync("input").toString();

const sample = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...

`;

const reverse = (str = "") => str.split("").reverse().join("");

const tiles = input
  .trim()
  .split("\n\n")
  .map((tile) => {
    const allLines = tile.split("\n");

    const id = parseInt(allLines[0].split(" ")[1].replace(":", ""));
    const lines = allLines.filter((_, idx) => idx != 0);
    const firstRow = lines[0];
    const lastRow = lines[lines.length - 1];
    const firstColumn = lines.map((l) => l[0]).join("");
    const lastColumn = lines.map((l) => l[l.length - 1]).join("");

    return {
      id,
      borders: [
        firstRow,
        reverse(firstRow),
        lastRow,
        reverse(lastRow),
        firstColumn,
        reverse(firstColumn),
        lastColumn,
        reverse(lastColumn),
      ],
    };
  });

console.log({
  part1: tiles.reduce((memo, tile) => {
    const adjacent = tiles
      .filter((t) => t.id !== tile.id)
      .reduce(
        (sum, t) =>
          t.borders.some((b) => tile.borders.includes(b)) ? sum + 1 : sum,
        0
      );

    return adjacent === 2 ? memo * tile.id : memo;
  }, 1),
});
