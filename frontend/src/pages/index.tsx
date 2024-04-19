import { lazy } from 'react';
import { Route, Routes } from "react-router-dom";

const Landing = lazy(() => import("./Landing"));
const Feed = lazy(() => import("./Feed"));
const Rating = lazy(() => import("./Rating"));
const Student = lazy(() => import("./Student"));
const Institution = lazy(() => import("./Institution"));

const Routing = () => {
  return (
    <Routes>
      <Route path="/" element={<Landing />} />
      <Route path="/feed" element={<Feed />} />
      <Route path="/rating/*" element={<Rating />} />
      <Route path="/student/*" element={<Student />} />
      <Route path="/inst/*" element={<Institution />} />
    </Routes>
  );
};

export default Routing;
