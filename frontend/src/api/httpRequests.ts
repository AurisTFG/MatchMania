import axiosClient from '../configs/axiosClient';

export const getRequest = async <T>({ url }: { url: string }): Promise<T> => {
  const response = await axiosClient.get<T>(url);

  return response;
};

export const postRequest = async <T>({
  url,
  body,
}: {
  url: string;
  body?: unknown;
}): Promise<T> => {
  const response = await axiosClient.post<T>(url, body);

  return response;
};

export const putRequest = async <T>({
  url,
  body,
}: {
  url: string;
  body?: unknown;
}): Promise<T> => {
  const response = await axiosClient.put<T>(url, body);

  return response;
};

export const deleteRequest = async <T>({
  url,
}: {
  url: string;
}): Promise<T> => {
  const response = await axiosClient.delete<T>(url);

  return response;
};

export const patchRequest = async <T>({
  url,
  body,
}: {
  url: string;
  body?: unknown;
}): Promise<T> => {
  const response = await axiosClient.patch<T>(url, body);

  return response;
};
