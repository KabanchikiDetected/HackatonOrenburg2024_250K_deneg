import { lazy, useLayoutEffect, useState } from "react";
import { jwtDecode } from "jwt-decode";

const UniversityProfile = lazy(
  () => import("features/components/UniversityProfile")
);
const UniversityReg = lazy(() => import("features/components/UniversityReg"));
const StudentProfile = lazy(() => import("features/components/StudentProfile"));
const StudentReg = lazy(() => import("features/components/StudentReg"));

const Profile = () => {
  const [render, setRender] = useState<string>("");
  const [token, _] = useState<string>(localStorage.getItem("token") as string);

  useLayoutEffect(() => {
    async function request() {
      //@ts-ignore
      let role = jwtDecode(token).role;

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
          setRender("user-reg");
        } else {
          setRender("user-profile");
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
};

export default Profile;
