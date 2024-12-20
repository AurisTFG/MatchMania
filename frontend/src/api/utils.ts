export function handleApiError(error: unknown): never {
  const message =
    (error as { response?: { data?: { error?: string } } })?.response?.data
      ?.error ||
    (error as Error).message ||
    "An unknown error occurred";
  throw new Error(message);
}
