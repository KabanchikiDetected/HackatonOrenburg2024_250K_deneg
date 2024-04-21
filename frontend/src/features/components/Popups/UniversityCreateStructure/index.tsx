/* @ts-ignore */
import classes from "./index.module.scss";

const UniversityEditStructure = ({ setIsOpen, faculties, setFaculties }: any) => {

  function addFaculty() {
    faculties.push({ faculty: "Факультет", id: 1, departments: [{ name: "Кафедра", id: 1 }] })
  }
  return (
    <div className={classes.wrapper}>
      <div className={classes.row + " " + classes.top}>
        <h2 className={classes.title}>Структура ОГУ</h2>
        <button className={classes.close} onClick={() => setIsOpen(false)}>
          <img src="/images/close.svg" alt="" />
        </button>
      </div>

      <div className={classes.row + " " + classes.bottom}>
        {/* <div className={classes.faculty}>
          <div className={classes.faculty}>Факультет</div>
          <div className={classes.department}>Кафедра</div>
        </div> */}
        {faculties.map((item: any, i1) => {
          return (
            <div className={classes.faculty}>
              <div className={classes.faculty_wrapper}>
              <input className={classes.faculty} />
              <img src="/images/close.svg" alt="" />
              </div>
              {item.departments.map((department, i2) => {
                return (
                  <div className={classes.department_wrapper}>
                    <input value={department} onChange={e => {
                      // setFaculties()
                    }} />

                  </div>

                );
              })}
              <button>Добавить кафедру</button>
            </div>
          );
        }
        )}
        <button
          onClick={addFaculty}
        >
          Добавить факультет
        </button>
      </div>
    </div>
  );
};

export default UniversityEditStructure;
