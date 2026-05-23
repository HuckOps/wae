declare namespace API {
  interface Restful<T> {
    code: number;
    message: string;
    data: T;
  }

  interface Pagination<T> {
    total: number;
    items: T[];
    page: number;
    page_size: number;
  }

  interface Service {
    id: number;
    created_at: string;
    updated_at: string;
    deleted_at: string | null;
    name: string;
    repo: string;
    domain: string;
    ref: string;
    creator: string;
    admins: string[];
    status: string;
    version: string;
    last_deploy: number;
    description: string;
    cluster: string;
  }

  interface Cluster {
    name: string;
    tags: string[];
  }
}
