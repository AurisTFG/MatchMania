import axiosClient from '../configs/axiosClient';

export const getRequest = async <T>({ url }: { url: string }): Promise<T> => {
  const response = await axiosClient.get<T>(url);

  return response as T;
};

export const postRequest = async <T>({
  url,
  body,
}: {
  url: string;
  body?: unknown;
}): Promise<T> => {
  const response = await axiosClient.post<T>(url, body);

  return response as T;
};

export const patchRequest = async <T>({
  url,
  body,
}: {
  url: string;
  body?: unknown;
}): Promise<T> => {
  const response = await axiosClient.patch<T>(url, body);

  return response as T;
};

export const deleteRequest = async ({
  url,
}: {
  url: string;
}): Promise<void> => {
  await axiosClient.delete(url);
};
