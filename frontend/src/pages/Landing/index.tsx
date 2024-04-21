import "./index.scss";
import { Link } from "react-router-dom";
import Footer from "features/components/Footer";

const Landing = () => {
  return (
    <main>
      <div className="top-bar">
        <div className="top-bar__logo">
          <img src="/images/landing_img_1.png" alt="" />
        </div>
        <div className="top-bar__panel panel">
          <div className="panel__text">
            Приветствуем вас в Lifecourse - приложении, которое упрощает процесс
            обучения и помогает вам достичь успеха. Наше приложение объединяет
            студентов, преподавателей, работодателей и родителей, обеспечивая
            удобное пространство для обучения и карьерного роста. Создайте свое
            впечатляющее резюме, найдите идеальную вакансию и достигните новых
            высот в обучении. Присоединяйтесь к нам сегодня и откройте для себя
            новые возможности!
          </div>
          <div className="panel__nav">
            <div className="link">
              <Link to="/login">ВУЗам</Link>
            </div>
            <div className="link">
              <Link to="/login">Студентам</Link>
            </div>
            <div className="link">
              <Link to="/login">Предприятиям</Link>
            </div>
          </div>
        </div>
      </div>
      {/* <Rating /> */}
      <Footer />
    </main>
  );
};

export default Landing;
