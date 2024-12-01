import app from "./app";
import config from "./config";
import { runTest } from "./tests/test";

app.listen(config.port, async () => {
  console.log(`🚀 ${config.name} ${config.version} 🚀`);
  console.log(
    `🚀 Listening on ${config.port} with NODE_ENV=${config.nodeEnv} 🚀`,
  );

  console.time("Execution Time");
  await runTest();
  console.timeEnd("Execution Time");
});
