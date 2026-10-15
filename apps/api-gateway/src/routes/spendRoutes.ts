
import { Router } from "express";
import { getSpendSummary } from "../services/spendClient";

const router = Router();

router.get("/summary", async (_req, res) => {
  const summary = await getSpendSummary();

  res.status(200).json(summary);
});

export default router;
