import { Fragment, ReactNode } from "react";
import { BrowserRouter } from "react-router-dom";

const Router = ({ children }: { children: ReactNode }) => {
  return (
    <Fragment>
      <BrowserRouter>{children}</BrowserRouter>
    </Fragment>
  );
};

export default Router;
