import { Fragment } from "react";
import { BrowserRouter } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import { Provider } from "react-redux";

import "react-toastify/dist/ReactToastify.css";

import store from "./redux/store";

import Routes from "./Routes";

function App() {
  return (
    <Fragment>
      <Provider store={store}>
        <BrowserRouter>
          <ToastContainer />
          {
            // TODO: Unit Tests are pending
            //TODO Error Boundary is pending
          }
          <Routes />
        </BrowserRouter>
      </Provider>
    </Fragment>
  );
}

export default App;
