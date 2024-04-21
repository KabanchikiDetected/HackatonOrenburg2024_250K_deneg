import { lazy, useLayoutEffect, useState } from "react";
import { jwtDecode } from "jwt-decode";
  // @ts-ignore

import Student from "pages/Student";
  // @ts-ignore

const UniversityProfile = lazy(
  () => import("features/components/UniversityProfile")
);
  // @ts-ignore
  const UniversityReg = lazy(() => import("features/components/UniversityReg"));
  // @ts-ignore
  const StudentProfile = lazy(() => import("features/components/StudentProfile"));
  // @ts-ignore
  const StudentReg = lazy(() => import("features/components/StudentReg"));

const Profile = () => {
  // @ts-ignore
  const [render, setRender] = useState<string>("");

  useLayoutEffect(() => {
    let token = localStorage.getItem("token");
  // @ts-ignore
  let role = jwtDecode(token).role;
    async function request() {

      if (role === "user") {
        let response = await fetch("/api/universities/my/education/", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        let ok = response.ok;

        if (!ok) {
          setRender("user-reg");
        } else {
          setRender("user-profile");
        }
      } else if (role === "deputy") {
        let response = await fetch("/api/universities/my/university/", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        let ok = response.ok;

        if (!ok) {
          setRender("uni-reg");
        } else {
          setRender("uni-profile");
        }
      }
    }

    request();
  }, []);

  if (render === "user-reg") {
    return <StudentReg />;
  } else if (render === "user-profile") {
    return <StudentProfile />;
  } else if (render === "uni-reg") {
    return <UniversityReg />;
  } else if (render === "uni-profile") {
    return <UniversityProfile />;
  }
  // return <UniversityProfile />
};

export default Profile;
