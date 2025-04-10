import { useQuery } from "@tanstack/react-query";
import { ENDPOINTS } from "../../constants/endpoints";
import { QUERY_KEYS } from "../../constants/queryKeys";
import { getRequest } from "../httpRequests";

export const useFetchUsers = () =>
  useQuery({
    queryKey: QUERY_KEYS.USERS.ALL,
    queryFn: () => getRequest({ url: ENDPOINTS.USERS.ROOT }),
  });
