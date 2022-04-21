export interface IConfig {
  appName: string;
  isProduction: boolean;
  baseURL: string;
  apiURL: string;
  apiURLV1: string;
  supportEmail: string;
}

export interface ICommonResponse {
  statusCode: number;
  message: string;
  data: any;
}

export interface ICommonListParams {
  limit: number;
  offset: number;

  search?: string;
}
