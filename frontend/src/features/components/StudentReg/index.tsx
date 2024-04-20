import { useRef, useState } from "react";
import "./index.scss";

const StudentReg = () => {
  const [data, setData] = useState({
    photoPreview: "",
    name: "",
    university: "",
    faculty: "",
    department: "",
    about: "",
  });
  const ref = useRef(null);

  function changeFile(e: any) {
    setData({ ...data, photoPreview: URL.createObjectURL(e.target.files[0]) });
  }

  function regStudent() {
    console.log(data);
  }

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
          <input
            id="university"
            value={data.university}
            onChange={(e) => setData({ ...data, university: e.target.value })}
            type="text"
          />
          <label htmlFor="faculty">
            <p>Факультет</p>
          </label>
          <input
            id="faculty"
            value={data.faculty}
            onChange={(e) => setData({ ...data, faculty: e.target.value })}
            type="text"
          />
          <label htmlFor="department">
            <p>Кафедра</p>
          </label>
          <input
            id="department"
            value={data.department}
            onChange={(e) => setData({ ...data, department: e.target.value })}
            type="text"
          />
        </div>
      </div>
      <textarea
        value={data.about}d
        onChange={(e) => setData({ ...data, about: e.target.value })}
        placeholder="О себе"
      />
      <br />
      <button onClick={regStudent}>Зарегистрироваться</button>
    </main>
  );
};

export default StudentReg;
