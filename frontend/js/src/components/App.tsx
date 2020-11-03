import React from "react";
import { connect } from "react-redux";

import { State } from "../state";

interface Props {
  message: string;
}

const App = ({ message }: Props) => <div>{message}</div>;

export default connect((state: State) => ({
  message: state.message,
}))(App);
