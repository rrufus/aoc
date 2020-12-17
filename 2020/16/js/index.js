const { readFileSync } = require("fs");
const input = readFileSync("../input").toString();
const toNumArray = (str = "") => str.split(",").map((v) => parseInt(v));
const myTicket = toNumArray(input.split("your ticket:\n")[1].split("\n")[0]);
const nearbyTickets = input
  .split("nearby tickets:\n")[1]
  .trim()
  .split("\n")
  .map(toNumArray);

const fieldsToRules = input
  .split("your ticket:\n")[0]
  .trim()
  .split("\n")
  .reduce((memo, line) => {
    const [field, ruleStr] = line.split(": ");
    return {
      ...memo,
      [field]: ruleStr.split(" or ").map((r) => {
        const [start, end] = r.split("-").map((v) => parseInt(v));
        return { start, end };
      }),
    };
  }, {});

const { errorRate, validTickets } = nearbyTickets
  .map((input = [0]) => {
    return input
      .map((i = 0) =>
        Object.values(fieldsToRules)
          .reduce((memo, rules) => [...memo, ...rules], [])
          .map((rule) => rule.start <= i && rule.end >= i)
          .includes(true)
          ? { valid: true, error: 0 }
          : { valid: false, error: i }
      )
      .reduce(
        (memo, current) => ({
          valid: memo.valid && current.valid,
          error: (memo.error += current.error),
        }),
        { valid: true, error: 0 }
      );
  })
  .reduce(
    (memo, current, idx) => ({
      errorRate: memo.errorRate + current.error,
      validTickets: current.valid
        ? [...memo.validTickets, nearbyTickets[idx]]
        : memo.validTickets,
    }),
    { errorRate: 0, validTickets: [] }
  );

const possibleFieldsAtIndex = validTickets
  .reduce(
    (memo, ticket) => ticket.map((v, idx) => (memo[idx] = [...memo[idx], v])),
    validTickets[0].map(() => [])
  )
  .reduce((memo, set = [0]) => {
    const matchingFields = Object.entries(fieldsToRules).reduce(
      (memo, [field, rules]) =>
        set.every((v) =>
          rules.map((rule) => rule.start <= v && v <= rule.end).includes(true)
        )
          ? [...memo, field]
          : memo,
      []
    );
    return [...memo, matchingFields];
  }, []);

let fieldToIndex = {};
while (Object.keys(fieldToIndex).length < Object.keys(fieldsToRules).length) {
  fieldToIndex = possibleFieldsAtIndex.reduce((memo, currList, idx) => {
    return currList.length === 1
      ? { ...memo, [currList[0]]: idx }
      : currList.filter((field) => memo[field] === undefined).length === 1
      ? {
          ...memo,
          [currList.filter((field) => memo[field] === undefined)[0]]: idx,
        }
      : memo;
  }, fieldToIndex);
}

const result = myTicket.reduce(
  (mem, v, idx) =>
    Object.entries(fieldToIndex)
      .filter(([field]) => field.startsWith("departure"))
      .map(([_, value]) => value)
      .includes(idx)
      ? mem * v
      : mem,
  1
);

console.log("Part 1: ", errorRate);
console.log("Part 2: ", result);
