import { useLayoutEffect, useRef, useState } from "react";
import "./index.scss";

const StudentReg = () => {
  const [universities, setUniversities] = useState([]);
  const [data, setData] = useState({
    photoPreview: "",
    name: "",
    birthday: "",
    university: "",
    faculty_id: "",
    department: "",
    about: "",
  });
  const ref = useRef(null);

  function changeFile(e: any) {
    setData({ ...data, photoPreview: URL.createObjectURL(e.target.files[0]) });
  }

  async function regStudent() {
    let response = await fetch("/api/students/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify(data),
    })

    if (response.ok) {
      response = await response.json();
    }

    console.log(response)

  }

  useLayoutEffect(() => {
    async function request() {
      const response = await fetch("/api/universities/university/");
      const data = await response.json();
      setUniversities(data);
    }

    request();
  }, [])

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
          <select name="university" id="">
            {universities.map((item) => {
              return (
                // @ts-ignore
                <option value={item.name} key={item.id}>
                  {/* @ts-ignore */}
                  {item.name}
                </option>
              );
            })}
          </select>
          <label htmlFor="faculty">
            <p>Факультет</p>
          </label>
          <select name="faculty" id="">
            <option value="ИМИТ">ИМИТ</option>
          </select>

          <label htmlFor="department">
            <p>Кафедра</p>
          </label>
          <select name="department" id="">
            <option value="ПОВТАЗ">ПОВТАЗ</option>
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
