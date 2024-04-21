/* @ts-ignore */
import classes from "./index.module.scss";

const UniversityStructure = ({ setIsOpen }: any) => {
  return (
    <div className={classes.wrapper}>
      <div className={classes.row + " " + classes.top}>
        <h2 className={classes.title}>Структура ОГУ</h2>
        <button className={classes.close} onClick={() => setIsOpen(false)}>
          <img src="/images/close.svg" alt="" />
        </button>
      </div>

      <div className={classes.row + " " + classes.bottom}>
        <div className="faculty">
          <div className="department">
            
          </div>
          <div className="department"></div>
          <div className="department"></div>
        </div>
      </div>
    </div>
  );
};

export default UniversityStructure;
