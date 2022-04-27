import { createStore, applyMiddleware, compose } from "redux";
import { createLogger } from "redux-logger";

import rootReducer from "./reducer";

//  Middlewares
let middlewares: any = [];
if (process.env.NODE_ENV !== "production") {
  const logger = createLogger();

  middlewares.push(logger);
}

const store = createStore(
  rootReducer,
  compose(applyMiddleware(...middlewares))
);

export default store;
