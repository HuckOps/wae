import { history } from '@umijs/max';
import toast from 'react-hot-toast';
import { extend, RequestOptionsInit } from 'umi-request';
import { getUser, removeUser } from './oidc';

const errorHandler = (error: any): false => {
  console.error(error);
  return false;
};

export interface ApiResponse<T = any> {
  code: number;
  message: string;
  data: T;
}

const extendedRequest = extend({
  prefix: '/api',
  errorHandler: errorHandler,
});

extendedRequest.interceptors.request.use((url, options) => {
  (async () => {
    const user = await getUser();
    if (user) {
      options.headers = {
        ...options.headers,
        Authorization: `Bearer ${user.access_token}`,
      };
    }
  })();
  return { url, options };
});

extendedRequest.interceptors.response.use(async (handler) => {
  try {
    const resp = await handler.json();
    if (handler.status >= 400) {
      if (handler.status === 401) {
        toast.error('Not login, redict to login page');
        await removeUser();
        history.push('/login');
      }
      toast.error(`Failed request to server, reason: ${resp.message}`);
      throw false;
    }
    if (resp.code !== 0) {
      toast.error(`Failed request to server, reason: ${resp.message}`);
      throw false;
    }
    return resp;
  } catch (error) {
    toast.error('Failed request to server, reason: ' + error);
    return false;
  }
});

export const baseRequest = {
  get: <T = any>(url: string, options?: RequestOptionsInit) =>
    extendedRequest.get<ApiResponse<T>>(url, options) as Promise<
      ApiResponse<T> | false
    >,
  post: <T = any>(url: string, options?: RequestOptionsInit) =>
    extendedRequest.post<ApiResponse<T>>(url, options) as Promise<
      ApiResponse<T> | false
    >,
  put: <T = any>(url: string, options?: RequestOptionsInit) =>
    extendedRequest.put<ApiResponse<T>>(url, options) as Promise<
      ApiResponse<T> | false
    >,
  delete: <T = any>(url: string, options?: RequestOptionsInit) =>
    extendedRequest.delete<ApiResponse<T>>(url, options) as Promise<
      ApiResponse<T> | false
    >,
};
