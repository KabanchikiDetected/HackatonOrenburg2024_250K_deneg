import { useRef, useState } from "react";
import "./index.scss";

const UniversityReg = () => {
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
            value={''}
            onChange={(e) => changeFile(e)}
            style={{ display: "none" }}
            ref={ref}
            type="file"
          />
          <button onClick={() => ref.current.click()}>Добавить фото</button>
        </div>

        <div className="info">
          <input
            value={data.name}
            onChange={(e) => setData({ ...data, name: e.target.value })}
            type="text"
            placeholder="Короткое название ВУЗа"
          />
          <input
            value={data.university}
            onChange={(e) => setData({ ...data, university: e.target.value })}
            type="text"
            placeholder=""
          />
          <input
            value={data.faculty}
            onChange={(e) => setData({ ...data, faculty: e.target.value })}
            type="text"
            placeholder="Город"
          />
        </div>
      </div>
      <textarea
        value={data.about}
        onChange={(e) => setData({ ...data, about: e.target.value })}
        placeholder="О себе"
      />
      <br />
      <button>Зарегистрироваться</button>
    </main>
  );
};

export default UniversityReg;
