import "./index.scss";
import { useLayoutEffect, useState } from "react";

interface IPost {
  id: number,
  title: string,
  content: string,
  author_id: string,
  created_at: number,
  likes: number,
  images: string[] | string,
  hashtags: string[] | string
}

const Feed = () => {
  const [token, _] = useState<string>(localStorage.getItem("token") as string);
  const [posts, setPosts] = useState<IPost[]>([])

  useLayoutEffect(() => {
    async function request() {
      await fetch("/api/news/news/feed//", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(response => response.json())
      .then(data => setPosts(data));
    }

    request();
  }, []);

  const getDate = (created_at: number) => {
    const date = new Date(created_at * 1000);
    const months = ['янв', 'фев', 'мар', 'апр', 'май', 'июн', 'июл', 'авг', 'сен', 'окт', 'ноя', 'дек'];
    
    const day = date.getDate().toString().padStart(2, '0');
    const month = months[date.getMonth()];
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${day} ${month} ${hours}:${minutes}`;
  }

  return (
    <main className="feed">
      <div className="posts">
        { posts.map(post => (
          <div className="post">
            <div className="post__row row">
              <div className="avatar">
                <img src="/images/user.png" alt="" />
              </div>
              <div className="about">
                <div className="about__col col">
                  <div className="name">
                    { post.author_id } <img src="/images/star.svg" alt="" /> <span>43</span>
                  </div>
                  <div className="date">{ getDate(post.created_at) }</div>
                </div>
                <div className="about__col col col-blue">
                  КУБГУ, Экономический, Инноватика, 215Б
                </div>
              </div>
            </div>
            <div className="post__row row">
              <img src="/images/post.png" alt="" />
              <img src="/images/post.png" alt="" />
            </div>
            <div className="post__col col">
              <strong className="title">
                { post.title }
              </strong>
              <div className="text">
                { post.content }
              </div>
            </div>
            <div className="post__row row row-space">
              <button className="like">
                <img src="/images/like.svg" alt="" /> <span> {post.likes} </span>
              </button>
              <p className="blue-bold">Творчество</p>
            </div>
          </div>
        )) }
        {/* <div className="post">
          <div className="post__row row">
            <div className="avatar">
              <img src="/images/user.png" alt="" />
            </div>
            <div className="about">
              <div className="about__col col">
                <div className="name">
                  Иван Иванов <img src="/images/star.svg" alt="" /> 43
                </div>
                <div className="date">9 апр в 21:31</div>
              </div>
              <div className="about__col col col-blue">
                КУБГУ, Экономический, Инноватика, 215Б
              </div>
            </div>
          </div>
          <div className="post__row row">
            <img src="/images/post.png" alt="" />
          </div>
          <div className="post__col col">
            <strong className="title">
              Чемпионат Приволжского федерального округа по гиревому спорту
            </strong>
            <div className="text">
              Всем привет
              <br />
              <br />
              Хочу поделиться радостью: Мы снова победили на
              хакатоне и заняли 2 место!!
              <br />
              <br />
              Нас от первого места отделил 1 бал,
              от чего грустно. поэтому мы обязательно вернемся туда за 1 местом,
              для нас это гештальт, который надо закрыть
            </div>
          </div>
          <div className="post__row row row-space">
            <button className="like">
              <img src="/images/like.svg" alt="" />
            </button>
            <p className="blue-bold">Творчество</p>
          </div>
        </div> */}
      </div>
    </main>
  );
};

export default Feed;
