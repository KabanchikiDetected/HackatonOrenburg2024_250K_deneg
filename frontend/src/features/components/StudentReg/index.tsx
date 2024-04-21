import { useEffect, useLayoutEffect, useRef, useState } from "react";
import "./index.scss";
import { useNavigate } from "react-router-dom";

const StudentReg = () => {
  const [universities, setUniversities] = useState([]);
  const [faculties, setFaculties] = useState([]);
  const [departments, setDepartments] = useState([]);
  const navigate = useNavigate()
  const [data, setData] = useState({
    photoPreview: "",
    name: "",
    birthday: "",
    university: "1",
    faculty_id: "1",
    department: "1",
    about: "",
  });

  const ref = useRef(null);

  function changeFile(e: any) {
    setData({ ...data, photoPreview: URL.createObjectURL(e.target.files[0]) });
  }

  async function regStudent() {
    let response = await fetch(
      `/api/universities/university/${data.university}/department/${data.faculty_id}/group/${data.department}/requests/`,
      {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("token")}`,
        }
      }
    )
    response = await fetch("/api/students/students/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify({
        birthday: `2000.01.01`,
        description: data.about,
        faculty_id: String(data.faculty_id),
        first_name: data.name.split(" ")[0],
        last_name: data.name.split(" ")[1],
      }),
    });

    if (response.ok) {
      response = await response.json();
      localStorage.setItem("user", JSON.stringify(response));
      navigate('/lk/feed')
    }
  }

  useLayoutEffect(() => {
    async function request() {
      let response = await fetch("/api/universities/university/");
      response = await response.json();
      setUniversities(response);
      setData({
        ...data,
        university: response[0].id,
      });
    }

    request();
  }, []);

  useEffect(() => {
    async function request() {
      if (data.university) {
        let response = await fetch(
          `/api/universities/university/${data.university}/department/`
        );
        response = await response.json();
        setFaculties(response);

        setData({
          ...data,
          faculty_id: response[0].id,
        });
      }
    }

    request();
  }, [data.university]);

  useEffect(() => {
    console.log(data);
  }, [data]);

  useEffect(() => {
    async function request() {
      if (data.university) {
        let response = await fetch(
          `/api/universities/university/${data.university}/department/${data.faculty_id}/group/`
        );
        response = await response.json();
        setDepartments(response);

        setData({
          ...data,
          department: response[0].id,
        });
      }
    }

    request();
  }, [data.faculty_id]);

  return (
    <main className="student-reg reg">
      <div className="reg__row">
        <div className="avatar">
          <img src={data.photoPreview} alt="" />
          <input
            value={""}
            onChange={(e) => changeFile(e)}
            style={{ display: "none" }}
            ref={ref}
            type="file"
          />
          {/* @ts-ignore */}

          <button onClick={() => ref.current.click()}>Добавить фото</button>
        </div>

        <div className="info">
          <label htmlFor="name">
            <p>Имя и фамилия</p>
          </label>

          <input
            id="name"
            value={data.name}
            onChange={(e) => setData({ ...data, name: e.target.value })}
            type="text"
          />
          <label htmlFor="university">
            <p>Университет</p>
          </label>
          <select
            name="university"
            id=""
            value={universities[data.university]}
            onChange={(e) => {
              const selectedUniversityId =
                e.target.selectedOptions[0].getAttribute("name");
              setData({ ...data, university: selectedUniversityId });
            }}
          >
            {universities.map((item) => {
              return (
                // @ts-ignore
                <option
                  value={item.name}
                  name={item.id}
                  key={item.name + item.id}
                >
                  {/* @ts-ignore */}
                  {item.name}
                </option>
              );
            })}
          </select>
          <label htmlFor="faculty">
            <p>Факультет</p>
          </label>
          <select name="faculty" value={faculties[data.faculty_id]} id="">
            {faculties.map((item) => {
              return (
                // @ts-ignore
                <option
                  value={item.name}
                  onChange={(e) => {
                    const selectedUniversityId =
                      e.target.selectedOptions[0].getAttribute("name");
                    setData({ ...data, faculty_id: selectedUniversityId });
                  }}
                  name={item.id}
                  key={item.name + item.id}
                >
                  {/* @ts-ignore */}
                  {item.name}
                </option>
              );
            })}
          </select>

          <label htmlFor="department">
            <p>Кафедра</p>
          </label>
          <select name="department" id="" value={departments[data.department]}>
            {departments.map((item) => {
              return (
                // @ts-ignore
                <option
                  onChange={(e) => {
                    const selectedUniversityId =
                      e.target.selectedOptions[0].getAttribute("name");
                    setData({ ...data, department: selectedUniversityId });
                  }}
                  value={item.name}
                  name={item.id}
                  key={item.name + item.id}
                >
                  {/* @ts-ignore */}
                  {item.name}
                </option>
              );
            })}
          </select>
        </div>
      </div>
      <textarea
        value={data.about}
        onChange={(e) => setData({ ...data, about: e.target.value })}
        placeholder="О себе"
      />
      <br />
      <button onClick={regStudent}>Зарегистрироваться</button>
    </main>
  );
};

export default StudentReg;
