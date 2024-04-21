import "./index.scss";
import { useEffect, useState } from "react";

interface IUser {
  id: string
  name: string,
  last_name: string,
  education: string,
  raiting: number,
}

const Rating = () => {
  const [users, setUsers] = useState<IUser[]>([])

  useEffect(() => {
    
  }, []);

  return (
    <main className="rating">
      <div className="rates">
        <div className="rate">
          <div className="rating">#9</div>
          <div className="avatar">
            <img src="/images/user.png" alt="" />
          </div>
          <div className="about">
            <div className="name">
              Иван Иванов <img src="/images/star.svg" alt="" /> 43
            </div>
            <div className="university">
              КУБГУ, Экономический, Инноватика, 215Б
            </div>
          </div>
          <div className="stars">
            <div className="stars__stars">
              <img src="/images/star.png" alt="" /> 43
            </div>
            <img src="/images/flag.svg" alt="" />
          </div>
        </div>
      </div>
    </main>
  );
};
export default Rating;
