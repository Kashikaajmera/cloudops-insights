
import express from "express";
import cors from "cors";
import helmet from "helmet";

import { env } from "./config/env";
import spendRoutes from "./routes/spendRoutes";
import { errorHandler } from "./middleware/errorHandler";

const app = express();

app.use(helmet());

app.use(
  cors({
    origin: "http://localhost:3000",
  })
);

app.use(express.json({ limit: "1mb" }));

app.get("/health", (_req, res) => {
  res.status(200).json({
    status: "ok",
    service: "api-gateway",
  });
});

app.use("/api/v1/spend", spendRoutes);

// Handle unmatched routes
app.use((_req, res) => {
  res.status(404).json({
    code: "NOT_FOUND",
    message: "Requested endpoint does not exist",
  });
});

// Error handler must be registered last
app.use(errorHandler);

app.listen(env.port, () => {
  console.log(`API Gateway running on port ${env.port}`);
});
