import { baseRequest } from '@/utils/baseRequest';

export const getServices = async (page?: number, pageSize?: number) => {
  const params: Record<string, number> = {};
  if (page) params.page = page;
  if (pageSize) params.page_size = pageSize;
  return await baseRequest.get<API.Pagination<API.Service>>('/v1/services', { params });
};

export const createService = async (data: {
  name: string;
  repo: string;
  domain: string;
  cluster: string;
  description?: string;
}) => {
  return await baseRequest.post<API.Service>('/v1/services', { data });
};

export const getClusters = async () => {
  return await baseRequest.get<API.Cluster[]>('/v1/clusters');
};