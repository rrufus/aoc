const { readFileSync } = require("fs");
const input = readFileSync("input").toString();

const reverse = (str = "") => str.split("").reverse().join("");

const flipHorizontal = (strArray = [""]) => strArray.map((str) => reverse(str));
const flipVertical = (strArray = [""]) => strArray.reverse();
const rotateClockwise = (strArray = [""]) =>
  strArray[0].split("").map((_, idx) =>
    strArray
      .map((s) => s[idx])
      .reverse()
      .join("")
  );
const toVariations = (strArray = [""]) => {
  const rotations = [
    strArray,
    rotateClockwise(strArray),
    rotateClockwise(rotateClockwise(strArray)),
    rotateClockwise(rotateClockwise(rotateClockwise(strArray))),
  ];
  return rotations
    .map((r) => [
      r,
      flipHorizontal(r),
      flipVertical(r),
      flipHorizontal(flipVertical(r)),
    ])
    .reduce((flatList, list) => [...flatList, ...list], []);
};

const getFirstRow = (lines = [""]) => lines[0];
const getLastRow = (lines = [""]) => lines[lines.length - 1];
const getFirstColumn = (lines = [""]) => lines.map((l) => l[0]).join("");
const getLastColumn = (lines = [""]) =>
  lines.map((l) => l[l.length - 1]).join("");

const getBorders = (lines = [""]) => {
  const firstRow = getFirstRow(lines);
  const lastRow = getLastRow(lines);
  const firstColumn = getFirstColumn(lines);
  const lastColumn = getLastColumn(lines);

  return [
    firstRow,
    reverse(firstRow),
    lastRow,
    reverse(lastRow),
    firstColumn,
    reverse(firstColumn),
    lastColumn,
    reverse(lastColumn),
  ];
};

const tiles = input
  .trim()
  .split("\n\n")
  .map((tile) => {
    const allLines = tile.split("\n");

    const id = parseInt(allLines[0].split(" ")[1].replace(":", ""));
    const lines = allLines.filter((_, idx) => idx != 0);

    return {
      id,
      lines,
      borders: getBorders(lines),
      variations: toVariations(lines),
    };
  });

console.log(
  "Part 1: " +
    tiles.reduce((memo, tile) => {
      const adjacent = tiles
        .filter((t) => t.id !== tile.id)
        .reduce(
          (sum, t) =>
            t.borders.some((b) => tile.borders.includes(b)) ? sum + 1 : sum,
          0
        );
      return adjacent === 2 ? memo * tile.id : memo;
    }, 1)
); // 27803643063307

const allMatches = tiles.map((t1) => {
  const matches = tiles
    .filter((t2) => t1.id !== t2.id)
    .filter((t2) => t1.borders.some((b) => t2.borders.includes(b)))
    .map((m) => m.id);
  return { ...t1, matches };
});

const firstCorner = allMatches.find((m) => m.matches.length === 2);

const firstCornerVariation = firstCorner.variations.find((v) => {
  const lastRow = getLastRow(v);
  const lastCol = getLastColumn(v);

  const m1 = allMatches.find((m) => m.id === firstCorner.matches[0]);
  const m2 = allMatches.find((m) => m.id === firstCorner.matches[1]);
  return (
    // (m1.borders.includes(lastRow) && m2.borders.includes(lastCol)) ||
    m1.borders.includes(lastCol) && m2.borders.includes(lastRow)
  );
});

let image = {
  placed: [firstCorner.id],
  image: Array.from(Array(tiles.length ** 0.5)).map(() =>
    Array.from(Array(tiles.length ** 0.5)).map(() => ({
      variation: [],
      id: null,
    }))
  ),
};
image.image[0][0] = { variation: firstCornerVariation, id: firstCorner.id };

const findVariationAndCoords = (image = [[]], id) => {
  let rowIdx = 0;
  while (rowIdx < image.length) {
    const colIdx = image[rowIdx].findIndex((tile) => tile.id === id);
    if (colIdx !== -1) {
      return { colIdx, rowIdx, variation: image[rowIdx][colIdx].variation };
    }
    rowIdx++;
  }
};

while (image.placed.length < allMatches.length) {
  image = allMatches.reduce((memo, current) => {
    if (memo.placed.includes(current.id)) {
      // already placed
      return memo;
    }
    const placedTiles = memo.placed.map((l) =>
      allMatches.find((m) => m.id === l)
    );

    if (!placedTiles.some((l) => l.matches.includes(current.id))) {
      // not adjacent to any tiles
      return memo;
    }

    const adjacentPlacedTiles = placedTiles
      .filter((l) => current.matches.includes(l.id))
      .map((l) => findVariationAndCoords(memo.image, l.id));

    const variation = current.variations
      .map((v) => {
        // since we are filling top left to bottom right we will only see these
        const firstCol = getFirstColumn(v);
        const firstRow = getFirstRow(v);

        const sameRow = adjacentPlacedTiles
          .map((a) => ({
            colIdx: a.colIdx,
            rowIdx: a.rowIdx,
            lastCol: getLastColumn(a.variation),
            lastRow: getLastRow(a.variation),
          }))
          .find((a) => a.lastCol === firstCol);

        const sameCol = adjacentPlacedTiles
          .map((a) => ({
            colIdx: a.colIdx,
            rowIdx: a.rowIdx,
            lastCol: getLastColumn(a.variation),
            lastRow: getLastRow(a.variation),
          }))
          .find((a) => a.lastRow === firstRow);

        if (sameRow) {
          return {
            variation: v,
            colIdx: sameRow.colIdx + 1,
            rowIdx: sameRow.rowIdx,
          };
        } else if (sameCol) {
          return {
            variation: v,
            colIdx: sameCol.colIdx,
            rowIdx: sameCol.rowIdx + 1,
          };
        }
      })
      .find((v) => !!v);

    if (variation) {
      const newImage = Array.from(Array(tiles.length ** 0.5)).map((_, rowIdx) =>
        Array.from(Array(tiles.length ** 0.5)).map(
          (_, colIdx) => memo.image[rowIdx][colIdx]
        )
      );

      newImage[variation.rowIdx][variation.colIdx] = {
        variation: variation.variation,
        id: current.id,
      };

      return {
        placed: [...memo.placed, current.id],
        image: newImage,
      };
    }
    return memo;
  }, image);
}

const removedBorders = image.image.map((row) => {
  return row.map((tile) => {
    return tile.variation
      .filter((_, idx) => idx !== 0 && idx !== tile.variation.length - 1)
      .map((line = "") => line.substring(1, line.length - 1));
  });
});

const toRowStrings = (input = [[""]]) =>
  input.reduce((memo, curr, idx) => {
    if (idx === 0) {
      return memo;
    }
    const newLines = [...memo];
    return newLines.map((l, idx) => `${l}${curr[idx]}`);
  }, input[0]);

const toImage = (rows = [[""]]) => {
  return rows.reduce((memo, current) => {
    return [...memo, ...toRowStrings(current)];
  }, []);
};

const finalImage = toImage(removedBorders);
const imageVariations = toVariations(finalImage);

const seaMonster = [
  "                  # ".split(""),
  "#    ##    ##    ###".split(""),
  " #  #  #  #  #  #   ".split(""),
];

seaMonsterPixels = 15;

const imagePixels = finalImage.reduce(
  (n, curr) =>
    n + curr.split("").reduce((t, curr) => (curr === "#" ? t + 1 : t), 0),
  0
);

imageVariations.find((v) => {
  const startRow = 0;
  const startCol = 0;
  const endRow = v.length - seaMonster.length;
  const endCol = v[0].length - seaMonster[0].length;

  let nSeaMonsters = 0;
  for (i = startRow; i <= endRow; i++) {
    for (j = startCol; j <= endCol; j++) {
      let seaMonsterFound = true;
      seaMonster.forEach((row, rowIdx) => {
        row.forEach((char, charIdx) => {
          if (char === "#") {
            if (v[i + rowIdx][j + charIdx] !== "#") {
              seaMonsterFound = false;
            }
          }
        });
      });
      if (seaMonsterFound) {
        nSeaMonsters++;
      }
    }
  }
  if (nSeaMonsters > 0) {
    // console.log(v);
    console.log("Part 2:", imagePixels - seaMonsterPixels * nSeaMonsters);
    return true;
  }
});
