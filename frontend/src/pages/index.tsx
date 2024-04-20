import Header from 'features/components/Header';
import Loader from 'features/components/Loader';
import { Fragment, lazy } from 'react';
import { Route, Routes } from "react-router-dom";

const Landing = lazy(() => import("./Landing"));
const Feed = lazy(() => import("./Feed"));
const Rating = lazy(() => import("./Rating"));
const Student = lazy(() => import("./Student"));
const University = lazy(() => import("./University"));
const Login = lazy(() => import("./Login"));
const Enterprise = lazy(() => import("./Enterprises"));
const Page404 = lazy(() => import("./Page404"));
const UniversityLK = lazy(() => import("./UniversityLK"));
const StudentLK = lazy(() => import("./StudentLK"));

const Routing = () => {
  return (
    <Fragment>
      <Header />
      <Routes>
        <Route path="/" element={<Loader children={<Landing />} />} />
        <Route path="/feed" element={<Loader children={<Feed />} />} />
        <Route path="/lk/student" element={<Loader children={<StudentLK />} />} />
        <Route path="/lk/university" element={<Loader children={<UniversityLK />} />} />
        <Route path="/rating" element={<Loader children={<Rating />} />} />
        <Route path="/students" element={<Loader children={<Student />} />} />
        <Route path="/universities" element={<Loader children={<University />} />} />
        <Route path="/enterprises" element={<Loader children={<Enterprise />} />} />
        <Route path="/login" element={<Loader children={<Login />} />} />
        <Route path="*" element={<Loader children={<Page404 />} />} />
      </Routes>
    </Fragment>
  );
};

export default Routing;
