import "./index.scss";

const Rating = () => {
  return (
    <div className="rating">
      <div className="rating__left">
        <img
          src="/images/star.png"
          alt=""
        />
        <p>ТОП студентов месяцев</p>
      </div>
      <div className="rating__right">
        {[
          1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
        ].map((item) => {
          return (
            <div className="rating__right__item" key={item}>
              <div className="rate">
                <img
                  src="/images/user.png"
                  alt=""
                />
                <p>Иван Иванов</p>
                <p>Оренбург</p>
                <p>ОГУ, ИМИТ</p>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Rating;
