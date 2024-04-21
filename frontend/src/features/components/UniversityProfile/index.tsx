import Modal from "features/components/Modal";
import "./index.scss";
import { useState } from "react";
import UniversityNewEvent from "../Popups/UniversityNewEvent";
import UniversityEditEvent from "../Popups/UniversityEditEvent";
import UniversityApplication from "../Popups/UniversityApplication";

const UniversityProfile = () => {
  const [isOpenCreate, setIsOpenCreate] = useState(false);
  const [isOpenEdit, setIsOpenEdit] = useState(false);
  const [isOpenStructure, setIsOpenStructure] = useState(false);
  const [isOpenJoin, setIsOpenJoin] = useState(false);

  return (
    <main className="profile">
      <Modal isOpen={isOpenCreate} setIsOpen={setIsOpenCreate}>
        <UniversityNewEvent setIsOpen={setIsOpenCreate} />
      </Modal>
      <Modal isOpen={isOpenEdit} setIsOpen={setIsOpenEdit}>
        <UniversityEditEvent setIsOpen={setIsOpenEdit} />
      </Modal>
      <Modal isOpen={isOpenStructure} setIsOpen={setIsOpenStructure}>
        <UniversityEditEvent setIsOpen={setIsOpenStructure} />
      </Modal>
      <Modal isOpen={isOpenJoin} setIsOpen={setIsOpenJoin}>
        <UniversityApplication setIsOpen={setIsOpenJoin} />
      </Modal>
      <div className="about">
        <div className="about__row row">
          <div className="row__logo">
            <img src="/images/user.png" alt="" />
          </div>
          <div className="row__about">
            <div className="short-info">
              <span>Оренбургский государственный университет | ОГУ</span>
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
        <div className="row__buttons">
          <button
            onClick={() => setIsOpenStructure(true)}
          >Структура</button>
          <button>Студенты</button>
        </div>
      </div>

      <button className="new-post" onClick={() => setIsOpenCreate(true)}>
        <div className="avatar" onClick={() => setIsOpenCreate(true)}>
          <img src="/images/user.png" alt="" />
        </div>
        <div className="wrapper" onClick={() => setIsOpenCreate(true)}>
          <p>Добавить мероприятие...</p>
          <img src="/images/plus.svg" alt="" />
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
              <button className="edit" onClick={() => setIsOpenEdit(true)}>
                <img src="/images/edit.svg" alt="" />
              </button>
            </div>
            <button className="button" onClick={() => setIsOpenJoin(true)}>
              Я участвовал
            </button>
          </div>
        </div>
      </div>
    </main>
  );
};

export default UniversityProfile;
