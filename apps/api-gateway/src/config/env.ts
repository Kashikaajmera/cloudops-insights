
import "dotenv/config";

const port = Number(process.env.PORT ?? "3001");

if (!Number.isInteger(port) || port < 1 || port > 65535) {
  throw new Error("PORT must be a valid TCP port");
}

const spendServiceUrl = process.env.SPEND_SERVICE_URL;

if (!spendServiceUrl) {
  throw new Error("SPEND_SERVICE_URL is required");
}

const parsedUrl = new URL(spendServiceUrl);

if (!["http:", "https:"].includes(parsedUrl.protocol)) {
  throw new Error("SPEND_SERVICE_URL must use HTTP or HTTPS");
}

export const env = {
  port,
  spendServiceUrl: parsedUrl.origin,
};
