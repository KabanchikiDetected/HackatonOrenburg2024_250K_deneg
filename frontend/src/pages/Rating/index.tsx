import "./index.scss";
import { useEffect, useState } from "react";

interface IUser {
  id: string,
  name: string,
  last_name: string,
  education?: string,
  rating?: number,
}

interface IRaiting {
  user_id: string,
  rating: number,
}

const Rating = () => {
  //@ts-ignore
  const [users, setUsers] = useState<IUser[]>([])

  useEffect(() => {
    async function reqeust() {
      const response = await fetch("/api/events/users/rating")
      const rating = await response.json()

      // const users = await data.map((raiting: IRaiting) => fetch(`/api/students/${raiting.user_id}`).then(response => response.json()))
      Promise.all(rating.map((raiting: IRaiting) => fetch(`/api/students/${raiting.user_id}`).then(response => response.json())))
      .then(data => data.map((userData, index) => ({ ...userData, rating: rating[index].rating })))

      
      //@ts-ignore
      .then((data: IUser) => {
        console.log(data)
        //@ts-ignore
        setUsers(data)
      })
      // setUsers(
      //   //@ts-ignore
      //   users.map((user, index) => ({
      //     ...user,
      //     raiting: data[index]
      //   }))
      // )
      // .then((data: IRaiting[]) => {
      //   Promise.all(data.map(raiting => fetch(`/api/students/${raiting.user_id}`).then(response => response.json())))
      //   .then(data => data.map((userData, index) => ({ ...userData[0], raiting: data[index].raiting })))
      //   .then(users => {
      //     setUsers(users)
      // })
    }

    reqeust();

  }, []);

  return (
    <main className="rating">
      <div className="rates">
        { users.map((user, index) => 
          <div className="rate">

            <div className="rating">#{index + 1}</div>
            <div className="avatar">
              <img src="/images/user.png" alt="" />
            </div>
            <div className="about">
              <div className="name">
                {user.name} {user.last_name}
              </div>
            </div>
            <div className="stars">
              <div className="stars__stars">
                <img src="/images/star.png" alt="" /> {user.rating}
              </div>
            </div>
          </div>
        ) }
  
        {/* <div className="rate">

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
        </div> */}
      </div>
    </main>
  );
};
export default Rating;
