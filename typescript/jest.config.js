/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  moduleNameMapper: {},
  collectCoverageFrom: ["**/*.{ts}", "!**/node_modules/**"],
};
