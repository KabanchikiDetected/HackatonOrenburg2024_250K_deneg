import { Link, Routes } from "react-router-dom";
import "./index.scss";

const StudentLK = () => {
  return (
    <div className="student">
      <aside className="sidebar-left">
        <div className="top">
          <p className="top__link link">
            <Link to={"/lk/student/feed"}>Лента</Link>
          </p>
          <p className="top__link link">
            <Link to={"/lk/student"}>Профиль</Link>
          </p>
          <p className="top__link link">
            <Link to={"/lk/student/rating"}>Рейтинг</Link>
          </p>
          <p className="top__link link">
            <Link to={"/lk/student/chats"}>Чаты</Link>
          </p>
          <p className="top__link link">
            <Link to={"/lk/student/vacancies"}>Вакансии</Link>
          </p>
        </div>
        <div className="bottom">
          <p className="bottom__link link">
            <Link to={"/support"}>Поддержка</Link>
          </p>
          <p className="bottom__link link">
            <Link to={"/market"}>Маркет</Link>
          </p>
        </div>
      </aside>
      <Routes>
        {/* <Route path="/lk/student" element={<StudentCard />} /> */}
      </Routes>
      <aside className="sidebar-right">
        <div className="top">
          <input type="text" />
          <button>
            <img src="/images/filters.svg" alt="" />
          </button>
        </div>
        <div className="bottom">

        </div>
      </aside>
    </div>
  );
};

export default StudentLK;
