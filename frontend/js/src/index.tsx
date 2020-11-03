import React from "react";
import { render } from "react-dom";
import { createStore } from "redux";
import { Provider } from "react-redux";

import { displayMessage } from "./actions";
import App from "./components/App";
import reducer from "./reducer";

const store = createStore(reducer as any);

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById("root")
);

store.dispatch(displayMessage("Hello, world!"));
