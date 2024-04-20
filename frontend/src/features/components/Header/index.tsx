import { useState } from "react";
import "./index.scss";
import { Link } from "react-router-dom";

interface IData {
  city: string;
}

const Header = () => {
  //@ts-ignore
  const [data, setData] = useState<IData>({ city: "Orenburg" });

  //@ts-ignore
  function getRoutingText(): string {
    return "landing";
  }

  return (
    <header className="header">
      <Link to={"/"}>
        <div className="header__logo">
          <img
            src="/images/logo.png/"
            alt=""
          />
        </div>
      </Link>
      {/* <div className="header__routing">{getRoutingText()}</div> */}
      {/* <div className="header__city">{data && data.city}</div> */}
      <div className="header__signup">
        <Link to="/login">Вход/Регистрация</Link>
      </div>
    </header>
  );
};

export default Header;
