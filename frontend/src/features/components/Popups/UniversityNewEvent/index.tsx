/* @ts-ignore */
import classes from "./index.module.scss";

const UniversityNewEvent = ({ setIsOpen }: any) => {
  return (
    <div className={classes.wrapper}>
      <div className={classes.row + " " + classes.top}>
        <h2 className={classes.title}>Новое мероприятие</h2>
        <button className={classes.close} onClick={() => setIsOpen(false)}>
          <img src="/images/close.svg" alt="" />
        </button>
      </div>

      <div className={classes.row + " " + classes.bottom}>
        <label htmlFor="title">Название</label>
        <input type="text" id="title" />
        <label htmlFor="text">Текст</label>
        <textarea id="text"></textarea>
        <button>Добавить картинку</button>
        <label htmlFor="direction">Категория</label>
        <select id="direction">
        {['Наука', "Спорт", "Творчество", "Волонтерство"].map((item) => {
                  return <option key={item} value={item}>{item}</option>;
                })}
        </select>
        <label htmlFor="points">Баллы за участие</label>
        <input type="text" id="points" />
        <button>Опубликовать</button>
      </div>
    </div>
  );
};

export default UniversityNewEvent;
