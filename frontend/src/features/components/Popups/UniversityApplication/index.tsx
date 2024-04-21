  // @ts-ignore
import classes from "./index.module.scss";

const StudentEditPost = ({ setIsOpen }: any) => {
  return (
    <div className={classes.wrapper}>
      <div className={classes.row + " " + classes.top}>
        <p></p>
        <button className={classes.close} onClick={() => setIsOpen(false)}>
          <img src="/images/close.svg" alt="" />
        </button>
      </div>

      <div className={classes.row + " " + classes.bottom}>
        <h1 className={classes.title}>Заявка отправлена</h1>
        <p className={classes.text}>
          После подтверждения вашего участия ВУЗом, вам будет начислен рейтинг
        </p>
      </div>
    </div>
  );
};

export default StudentEditPost;
