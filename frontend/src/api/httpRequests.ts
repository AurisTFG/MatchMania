import axiosClient from 'configs/axiosClient';

export const getRequest = async <T>({ url }: { url: string }): Promise<T> => {
  const response = await axiosClient.get<T>(url);

  return response.data;
};

export const postRequest = async ({
  url,
  body,
}: {
  url: string;
  body?: unknown;
}): Promise<void> => {
  await axiosClient.post(url, body);
};

export const patchRequest = async ({
  url,
  body,
}: {
  url: string;
  body?: unknown;
}): Promise<void> => {
  await axiosClient.patch(url, body);
};

export const deleteRequest = async ({
  url,
}: {
  url: string;
}): Promise<void> => {
  await axiosClient.delete(url);
};
