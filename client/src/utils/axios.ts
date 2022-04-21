import axios, { AxiosError, AxiosRequestHeaders } from "axios";

import { MESSAGES, STATUS_CODES } from "../config/constants";

import { errorToast, warningToast } from "./toast";

interface IConfig {
  contentType?: string;
  skipErrorMsg?: boolean;
}

const errorResHandler = (error: AxiosError, config: IConfig) => {
  console.error(error);

  const statusCode = error.response?.status || error.code || STATUS_CODES.ERROR;
  const errorMsg =
    error.response?.data?.message || error.message || MESSAGES.errorMsg;

  if (!config.skipErrorMsg) {
    if (statusCode === STATUS_CODES.ACTION_FAILED) {
      warningToast(errorMsg);
    } else {
      errorToast(errorMsg);
    }
  }

  throw error;
};

export const makeGetRequest = (
  url: string,
  params: any = {},
  config: IConfig = {
    contentType: "application/x-www-form-urlencoded",
    skipErrorMsg: false,
  }
) => {
  const headers: AxiosRequestHeaders = {
    "Content-Type": config.contentType as string,
  };

  return axios
    .get(url, {
      params,
      headers,
    })
    .catch((error) => errorResHandler(error, config));
};

export const makePostRequest = (
  url: string,
  data: any = {},
  config: IConfig = {
    contentType: "application/json",
    skipErrorMsg: false,
  }
) => {
  const headers: AxiosRequestHeaders = {
    "Content-Type": config.contentType as string,
  };

  return axios
    .post(url, data, {
      headers,
    })
    .catch((error) => errorResHandler(error, config));
};

export const makePutRequest = (
  url: string,
  data: any = {},
  config: IConfig = {
    contentType: "application/json",
    skipErrorMsg: false,
  }
) => {
  const headers: AxiosRequestHeaders = {
    "Content-Type": config.contentType as string,
  };

  return axios
    .put(url, data, {
      headers,
    })
    .catch((error) => errorResHandler(error, config));
};
