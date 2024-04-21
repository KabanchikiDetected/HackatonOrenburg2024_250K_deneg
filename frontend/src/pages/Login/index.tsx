import { useState } from "react";
import "./index.scss";
import { useNavigate } from "react-router-dom";

interface IData {
  email: string | "";
  password: string | "";
  repeat_password: string | "";
}

const Login = () => {
  const [register, setRegister] = useState<boolean>(false);
  const navigate = useNavigate();
  const [data, setData] = useState<IData>({
    email: "",
    password: "",
    repeat_password: "",
  });
  async function handleLoginClick() {
    if (data.email && data.password) {
      let response = await fetch("/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: data.email,
          password: data.password,
        }),
      });
      response = await response.json();
      // @ts-ignore
      if (response.token) {
        // @ts-ignore
        localStorage.setItem("token", response.token);
        navigate("/lk/profile");
      } else {
        alert("Неверные данные");
      }
    }
  }

  async function handleRegisterClick() {
    if (data.email && data.password && data.repeat_password) {
      const response = await fetch("/api/auth/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        setRegister(false);
      }
    }
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
              value={data.repeat_password}
              onChange={(e) =>
                setData({ ...data, repeat_password: e.target.value })
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

            // @ts-ignore
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
