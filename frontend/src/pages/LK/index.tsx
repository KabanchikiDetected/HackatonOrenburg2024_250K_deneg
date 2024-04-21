import { Link, Route, Routes, useLocation } from "react-router-dom";
import "./index.scss";
import { lazy, useState } from "react";
import Loader from "features/components/Loader";
import cities from "features/utils/cities";
import { NavLink } from "react-router-dom";

const Profile = lazy(() => import("pages/Profile"));
const Feed = lazy(() => import("pages/Feed"));
const Rating = lazy(() => import("pages/Rating"));
const Register = lazy(() => import("pages/Register"));

interface IFilter {
  city: string;
  university: string;
  rating: string;
  education: string;
  faculty: string;
}

const StudentLK = () => {
  const [filters, setFilters] = useState<IFilter>({
    city: "Оренбург",
    university: "МГТУ им. Н.Э. Барабина",
    rating: "4",
    education: "Бакалавриат",
    faculty: "Информационные системы и технологии",
  });
  const location = useLocation();
  const [isFilters, setIsFilters] = useState<boolean>();

  return (
    <div className="student">
      <aside className="sidebar-left">
        <div className="top">
          <p className="top__link link">
            <NavLink to={"/lk/feed"}>Лента</NavLink>
          </p>
          <p className="top__link link">
            <NavLink to={"/lk/profile"}>Профиль</NavLink>
          </p>
          <p className="top__link link">
            <NavLink to={"/lk/rating"}>Рейтинг</NavLink>
          </p>
          <p className="top__link link">
            <NavLink to={"/lk/chats"}>Чаты</NavLink>
          </p>
          <p className="top__link link">
            <NavLink to={"/lk/vacancies"}>Вакансии</NavLink>
          </p>
          <p className="top__link link">
            <NavLink to={"/lk/settings"}>Настройки</NavLink>
          </p>
          <p className="top__link link">
            <NavLink to={"/lk/applications"}>Заявки</NavLink>
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
        <Route path="/profile" element={<Loader children={<Profile />} />} />
        <Route path="/feed" element={<Loader children={<Feed />} />} />
        <Route path="/rating" element={<Loader children={<Rating />} />} />
        <Route path="/register" element={<Loader children={<Register />} />} />
      </Routes>

      {location.pathname === "/lk/feed" ||
      location.pathname === "/lk/rating" ? (
        <aside className="sidebar-right">
          <div className="top">
            {
              location.pathname === "/lk/feed"
              && (
                <input placeholder="Поиск..." type="text" />
              )
            }
            
            <button
              onClick={() => setIsFilters(!isFilters)}
            >
              <img src="/images/filters.svg" alt="" />
            </button>
          </div>
          {
            isFilters
            && (
              <div className="bottom">
            <p>
              <p>Город</p>
              <select
                value={filters.city}
                onChange={(event) => {
                  setFilters({
                    ...filters,
                    city: event.target.value,
                  });
                }}
              >
                {cities.map((item, id) => {
                  return (
                    <option key={item + id} value={item}>
                      {item}
                    </option>
                  );
                })}
              </select>
            </p>
            <p>
              <p>ВУЗ</p>
              <select
                value={filters.university}
                onChange={(event) => {
                  setFilters({
                    ...filters,
                    city: event.target.value,
                  });
                }}
              >
                {/* {cities.map((item) => {
                  return <option key={item+2} value={item}>{item}</option>;
                })} */}
              </select>
            </p>
            <p>
              <p>Факультет</p>
              <select
                value={filters.faculty}
                onChange={(event) => {
                  setFilters({
                    ...filters,
                    city: event.target.value,
                  });
                }}
              >
                {/* {cities.map((item) => {
                  return <option key={item+3} value={item}>{item}</option>;
                })} */}
              </select>
            </p>
            <p>
              <p>Поток</p>
              <select
                value={filters.city}
                onChange={(event) => {
                  setFilters({
                    ...filters,
                    city: event.target.value,
                  });
                }}
              >
                {/* {cities.map((item) => {
                  return <option key={item+4} value={item}>{item}</option>;
                })} */}
              </select>
            </p>
            <p>
              <p>Кафедра</p>
              <select
                value={filters.city}
                onChange={(event) => {
                  setFilters({
                    ...filters,
                    city: event.target.value,
                  });
                }}
              >
                {/* {cities.map((item) => {
                  return <option key={item+5} value={item}>{item}</option>;
                })} */}
              </select>
            </p>
            <p>
              <p>Группа</p>
              <select
                value={filters.city}
                onChange={(event) => {
                  setFilters({
                    ...filters,
                    city: event.target.value,
                  });
                }}
              >
                {/* {cities.map((item) => {
                  return <option key={item+6} value={item}>{item}</option>;
                })} */}
              </select>
            </p>
            <p>
              {location.pathname === "/lk/rating" ? (
                <>
                  <p>Период</p>
                  <select
                    value={filters.city}
                    onChange={(event) => {
                      setFilters({
                        ...filters,
                        city: event.target.value,
                      });
                    }}
                  >
                    {/* {cities.map((item) => {
                  return <option key={item+7} value={item}>{item}</option>;
                })} */}
                  </select>
                </>
              ) : (
                <>
                  <p>Достижение</p>
                  <select
                    value={filters.city}
                    onChange={(event) => {
                      setFilters({
                        ...filters,
                        city: event.target.value,
                      });
                    }}
                  >
                    {['Наука', "Спорт", "Творчество", "Волонтерство"].map((item) => {
                  return <option key={item} value={item}>{item}</option>;
                })}
                  </select>
                </>
              )}
            </p>
          </div>
            )
          }
        </aside>
      ) : (
        <></>
      )}
    </div>
  );
};

export default StudentLK;
