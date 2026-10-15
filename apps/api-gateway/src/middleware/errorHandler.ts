
import type { ErrorRequestHandler } from "express";
import { AppError } from "../errors/AppError";

export const errorHandler: ErrorRequestHandler = (
  err,
  _req,
  res,
  _next
) => {
  if (err instanceof AppError) {
    res.status(err.statusCode).json({
      code: err.code,
      message: err.message,
    });
    return;
  }

  console.error("Unexpected error:", err);

  res.status(500).json({
    code: "INTERNAL_SERVER_ERROR",
    message: "An unexpected error occurred",
  });
};
