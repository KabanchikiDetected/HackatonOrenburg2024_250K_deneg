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
    city: "",
    short_name: "",
    long_name: "",
  });
  const ref = useRef(null);

  function changeFile(e: any) {
    setData({ ...data, photoPreview: URL.createObjectURL(e.target.files[0]) });
  }

  async function regUniversity() {
    let response = await fetch("/api/universities/university/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify(data),
    })

    if (!response.ok) {
      response = await response.json();
    }

    console.log(response)
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
  {/* @ts-ignore */}
          <button className="add-photo" onClick={() => ref.current.click()}>
            Добавить фото
          </button>
        </div>

        <div className="info">
          <label htmlFor="short-name">
            <p>Короткое название ВУЗа</p>
          </label>
          <input
            id="short-name"
            value={data.short_name}
            onChange={(e) => setData({ ...data, short_name: e.target.value })}
            type="text"
          />
          <label htmlFor="university">
            <p>Полное название ВУЗа</p>
          </label>
          <input
            id="university"
            value={data.long_name}
            onChange={(e) => setData({ ...data, long_name: e.target.value })}
            type="text"
          />
          <label htmlFor="faculty">
            <p>Город</p>
          </label>
          <input
            id="city"
            value={data.city}
            onChange={(e) => setData({ ...data, city: e.target.value })}
            type="text"
          />

          <button className="add-struct">
            <p>Добавить структуру</p>
          </button>
        </div>
      </div>
      <label htmlFor="about">О ВУЗе</label>
      <textarea
        value={data.about}
        onChange={(e) => setData({ ...data, about: e.target.value })}
        id="about"
      />
      <br />
      <button onClick={regUniversity} className="submit">
        Создать
      </button>
    </main>
  );
};

export default StudentReg;
