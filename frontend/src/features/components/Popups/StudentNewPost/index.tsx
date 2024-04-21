  // @ts-ignore
import classes from "./index.module.scss";

const StudentNewPost = ({ setIsOpen }: any) => {
  return (
    <div className={classes.wrapper}>
      <div className={classes.row + " " + classes.top}>
        <h2 className={classes.title}>Новый пост</h2>
        <button className={classes.close} onClick={() => setIsOpen(false)}>
          <img src="/images/close.svg" alt="" />
        </button>
      </div>

      <div className={classes.row + " " + classes.bottom}>
        <label htmlFor="title">Название</label>
        <input type="text" id="title" />
        <label htmlFor="text">Текст</label>
        <textarea id="text"></textarea>
        <button>Прикрепить картинку</button>
        <label htmlFor="direction">Категория</label>
        <select id="direction">
        {['Наука', "Спорт", "Творчество", "Волонтерство"].map((item) => {
                  return <option key={item} value={item}>{item}</option>;
                })}
        </select>
        <button>Опубликовать</button>
      </div>
    </div>
  );
};

export default StudentNewPost;
