import { readFileSync } from "fs";
import path from "path";

export async function runTest() {
  const jsonPath = path.resolve(__dirname, "./input.txt");
  const contents = readFileSync(jsonPath, "utf8").trim().split("\n");

  console.log(`result: => `);
}
