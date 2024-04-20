import { useState } from "react";
import "./index.scss";

interface IData {
  email: string | "";
  password: string | "";
  repeatPassword: string | "";
}

const Login = () => {
  const [register, setRegister] = useState<boolean>(false);
  const [data, setData] = useState<IData>({
    email: "",
    password: "",
    repeatPassword: "",
  });
  function handleLoginClick() {
    console.log("login");
  }

  function handleRegisterClick() {
    console.log("register");
  }

  return (
    <main className="login">
      <div className="login__wrapper">
        <h1>{register ? "Регистрация" : "Вход"}</h1>
        {register ? (
          <div className="wrapper">
            <label htmlFor="email" className="login__label">
              Введите email
            </label>
            <input
              value={data.email}
              onChange={(e) => setData({ ...data, email: e.target.value })}
              id="email"
              type="email"
              className="login__input"
            ></input>
            <label htmlFor="password" className="login__label">
              Введите пароль
            </label>
            <input
              value={data.password}
              onChange={(e) => setData({ ...data, password: e.target.value })}
              id="password"
              type="password"
              className="login__input"
            ></input>
            <label htmlFor="repeat-password" className="login__label">
              Введите пароль
            </label>
            <input
              value={data.repeatPassword}
              onChange={(e) =>
                setData({ ...data, repeatPassword: e.target.value })
              }
              id="repeat-password"
              type="password"
              className="login__input"
            ></input>
            <button onClick={handleRegisterClick} className="login__button">
              Зарегистрироваться
            </button>
          </div>
        ) : (
          <div className="wrapper">
            <label htmlFor="email" className="login__label">
              Введите email
            </label>
            <input
              value={data.email}
              onChange={(e) => setData({ ...data, email: e.target.value })}
              id="email"
              type="email"
              className="login__input"
            ></input>
            <label htmlFor="password" className="login__label">
              Введите пароль
            </label>
            <input
              value={data.password}
              onChange={(e) => setData({ ...data, password: e.target.value })}
              id="password"
              type="password"
              className="login__input"
            ></input>
            <button onClick={handleLoginClick} className="login__button">
              Войти
            </button>
          </div>
        )}
        <p>{!register ? "У вас нет аккаунта?" : "У вас уже есть аккаунт?"}</p>
        <button
          className="login__switch"
          onClick={() => {
            setRegister(!register);
            setData({ email: "", password: "", repeatPassword: "" });
          }}
        >
          {!register ? "Регистрация" : "Вход"}
        </button>
      </div>
    </main>
  );
};

export default Login;
