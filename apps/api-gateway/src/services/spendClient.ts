
import { AppError } from "../errors/AppError";
import { env } from "../config/env";

export interface SpendSummary {
  currency: string;
  totalCost: string;
  resourceCount: number;
}

export async function getSpendSummary(): Promise<SpendSummary> {
const baseUrl = env.spendServiceUrl;
  if (!baseUrl) {
    throw new AppError(
      500,
      "SERVICE_CONFIGURATION_ERROR",
      "Spend service URL is not configured"
    );
  }

  let response: Response;

  try {
    response = await fetch(
      `${baseUrl}/api/v1/spend/summary`,
      {
        signal: AbortSignal.timeout(5000),
      }
    );
  } catch (error) {
    console.error("Spend service connection failed:", error);

    throw new AppError(
      502,
      "SPEND_SERVICE_UNAVAILABLE",
      "Unable to communicate with spend service"
    );
  }

  if (!response.ok) {
    console.error(
      `Spend service returned HTTP ${response.status}`
    );

    throw new AppError(
      502,
      "SPEND_SERVICE_ERROR",
      "Spend service returned an unsuccessful response"
    );
  }

  try {
    return (await response.json()) as SpendSummary;
  } catch {
    throw new AppError(
      502,
      "INVALID_UPSTREAM_RESPONSE",
      "Spend service returned invalid JSON"
    );
  }
}
