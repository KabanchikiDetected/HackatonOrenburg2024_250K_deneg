  // @ts-ignore
  import classes from "./index.module.scss";

const StudentEditPost = ({ setIsOpen }: any) => {
  return (
    <div className={classes.wrapper}>
      <div className={classes.row + " " + classes.top}>
        <h2 className={classes.title}>Пост от 9 апр в 21:31</h2>
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
        <div className={classes.buttons}>
          <button className={classes.red}>Удалить</button>
          <button>Опубликовать</button>
        </div>
      </div>
    </div>
  );
};

export default StudentEditPost;
