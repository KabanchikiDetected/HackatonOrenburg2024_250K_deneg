import { Context } from "./Context";
import { UpdateContext } from "./UpdateContext";
import { Fragment, ReactNode, useState } from "react";

const ContextProvider = ({ children }: { children: ReactNode }) => {
  const [context, setContext] = useState<{}>();
  return (
    <Fragment>
      <Context.Provider value={[context]}>
        <UpdateContext.Provider value={[setContext]}>
          {children}
        </UpdateContext.Provider>
      </Context.Provider>
    </Fragment>
  );
};

export default ContextProvider;

export { Context, UpdateContext };
