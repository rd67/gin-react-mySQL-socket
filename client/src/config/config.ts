import * as interfaces from "./interfaces";

const isProduction = process.env.NODE_ENV === "production";

const baseURL = process.env.REACT_APP_BASE_URL as string;

const apiURL = process.env.REACT_APP_API_URL as string;

const config: interfaces.IConfig = {
  appName: "Apsis",
  isProduction,
  baseURL,
  apiURL,
  apiURLV1: `${apiURL}/v1`,
  supportEmail: "rohitdalal67@gmail.com",
};

export default config;
