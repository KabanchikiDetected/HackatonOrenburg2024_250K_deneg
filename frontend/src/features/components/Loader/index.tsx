import "./Loader.scss";
import { Suspense, ReactNode } from "react";
import { Circles } from "react-loader-spinner";

const Loader = ({ children }: { children: ReactNode }) => {
  return (
    <Suspense
      fallback={
        <Circles wrapperClass="loading" color="#4e5dbf" />
      }
    >
      {children}
    </Suspense>
  );
};

export default Loader;
