import { useState } from "react";
import "./index.scss";
import Modal from "features/components/Modal";
import StudentNewPost from "features/components/Popups/StudentNewPost";
import StudentEditPost from "../Popups/StudentEditPost";
import RatingHistory from "../Popups/RatingHistory";

const StudentProfile = () => {
  const [isOpenCreate, setIsOpenCreate] = useState(false);
  const [isOpenEdit, setIsOpenEdit] = useState(false);
  const [isOpenRating, setIsOpenRating] = useState(false);

  return (
    <main className="profile">
      <Modal isOpen={isOpenCreate} setIsOpen={setIsOpenCreate}>
        <StudentNewPost setIsOpen={setIsOpenCreate} />
      </Modal>
      <Modal isOpen={isOpenEdit} setIsOpen={setIsOpenEdit}>
        <StudentEditPost setIsOpen={setIsOpenEdit} />
      </Modal>
      <Modal isOpen={isOpenRating} setIsOpen={setIsOpenRating}>
        <RatingHistory setIsOpen={setIsOpenRating} />
      </Modal>
      <div className="about">
        <div className="about__row row">
          <div className="row__logo">
            <img src="/images/user.png" alt="" />
          </div>
          <div className="row__about">
            <div className="short-info">
              <span>Иван Головачев</span> <img src="/images/star.svg" alt="" />{" "}
              <span className="stars">43</span>
            </div>
            <div className="university">
              КУБГУ, Экономический, Инноватика, 215Б, #5
            </div>
            <div className="description">
              Такой то такойто и тдТакой то такойто и тдТакой то такойто и
              тдТакой то такойто и тдТакой то такойто и тдТакой то такойто и
              тдТакой то такойто и тдТакой то такойто и тдТакой то такойто и
              тдТакой то такойто и тдТакой то такойто и тдТакой то такойто и
              тдТакой то такойто и тд
            </div>
          </div>
        </div>
        <div className="rating__row"
        onClick={() => setIsOpenRating(true)}
        >
          <div className="rate">
            <div className="level">1 уровень</div>
            <div className="img">
              <img src="/images/science.svg" alt="" />
            </div>
            <div className="stars">
              <img src="/images/star.svg" alt="" />{" "}
              <span className="stars-label">10</span>
            </div>
          </div>
          <div className="rate">
            <div className="level">1 уровень</div>
            <div className="img">
              <img src="/images/sport.svg" alt="" />
            </div>
            <div className="stars">
              <img src="/images/star.svg" alt="" />{" "}
              <span className="stars-label">10</span>
            </div>
          </div>
          <div className="rate">
            <div className="level">1 уровень</div>
            <div className="img">
              <img src="/images/art.svg" alt="" />
            </div>
            <div className="stars">
              <img src="/images/star.svg" alt="" />{" "}
              <span className="stars-label">10</span>
            </div>
          </div>
          <div className="rate">
            <div className="level">1 уровень</div>
            <div className="img">
              <img src="/images/volunteer.svg" alt="" />
            </div>
            <div className="stars">
              <img src="/images/star.svg" alt="" />{" "}
              <span className="stars-label">10</span>
            </div>
          </div>
        </div>
      </div>

      <button
        className="new-post"
        onClick={() => {
          setIsOpenCreate(true);
        }}
      >
        <div
          className="avatar"
          onClick={() => {
            setIsOpenCreate(true);
          }}
        >
          <img src="/images/user.png" alt="" />
        </div>
        <div className="wrapper">
          <p>Создать пост...</p>
          <img
            src="/images/note.svg"
            onClick={() => {
              setIsOpenCreate(true);
            }}
            alt=""
          />
        </div>
      </button>

      <div className="posts">
        <div className="post">
          <div className="post__row row">
            <div className="avatar">
              <img src="/images/user.png" alt="" />
            </div>
            <div className="about">
              <div className="about__col col">
                <div className="name">
                  Иван Иванов <img src="/images/star.svg" alt="" />{" "}
                  <span>43</span>
                </div>
                <div className="date">9 апр в 21:31</div>
              </div>
              <div className="about__col col col-blue">
                КУБГУ, Экономический, Инноватика, 215Б
              </div>
            </div>
          </div>
          <div className="post__row row">
            <img src="/images/post.png" alt="" />
          </div>
          <div className="post__col col">
            <strong className="title">
              Чемпионат Приволжского федерального округа по гиревому спорту
            </strong>
            <div className="text">
              Всем привет Хочу поделиться радостью: Мы снова победили на
              хакатоне и заняли 2 место!! Нас от первого места отделил 1 бал, от
              чего грустно. поэтому мы обязательно вернемся туда за 1 местом,
              для нас это гештальт, который надо закрыть
            </div>
          </div>
          <div className="post__row row row-space">
            <div className="wrapper">
              <button className="like">
                <img src="/images/like.svg" alt="" />
              </button>
              <button className="edit"
              onClick={() => {
                setIsOpenEdit(true);
              }}
              >
                <img src="/images/edit.svg" alt="" />
              </button>
            </div>
            <p className="blue-bold">Творчество</p>
          </div>
        </div>
      </div>
    </main>
  );
};

export default StudentProfile;
