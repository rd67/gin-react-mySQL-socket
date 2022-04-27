import { combineReducers } from "redux";

const rootReducer = combineReducers({});

export type IReduxState = ReturnType<typeof rootReducer>;

export default rootReducer;
