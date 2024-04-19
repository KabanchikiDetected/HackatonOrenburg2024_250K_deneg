import { Fragment } from "react";
import "./index.scss";
import Routing from "pages";
import Router from "./Router/index.tsx";
import ContextProvider from "./contexts/index.tsx";

const App = () => {
  return (
    <Fragment>
      <ContextProvider>
      <Router>
        <Routing />
      </Router>
      </ContextProvider>
    </Fragment>
  );
};

export default App;
