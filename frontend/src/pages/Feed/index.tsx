import { Link } from "react-router-dom";
import "./index.scss";
import { useEffect, useState } from "react";

interface IUser {
  id: string
  name: string,
  last_name: string,
  description: string,
  birthday: "2007-04-14T00:00:00Z",
  faculty_id: string,
  photo: string,
  role: string,
  education?: string,
}

interface IPost {
  id: number,
  title: string,
  content: string,
  author_id: string,
  created_at: number,
  likes: number,
  images: string,
  hashtags: string,
  parsed_hashtags?: string[],
  author?: IUser,
  is_liked?: boolean,
}

const Feed = () => {
  const [token, _] = useState<string>(localStorage.getItem("token") as string);
  const [posts, setPosts] = useState<IPost[]>([])

  useEffect(() => {
    loadPosts();

  }, []);

  const loadPosts = () => {
    fetch('/api/news/feed/')
    .then(response => response.json())
    .then(data => {
      Promise.all(
        data.map((post: IPost) =>
            fetch(`/api/students/${post.author_id}`)
              .then(response => response.json())
              .then(authorData => ({ ...post, author: {
                ...authorData,
                education: "No University",
              }, parsed_hashtags: getHeshtags(post), is_liked: isLiked(post.id) }))
        )
      )
      .then(postsData => {
        Promise.all(
          postsData.map(post =>
            getEducation(post.author_id)
              .then(education => ({ ...post, author: {
                ...post.author,
                education
              }}))
          )
        )
        .then(postsDataWithEducation => {
          setPosts(postsDataWithEducation);
        });
      });
    })
    .catch(error => console.error('Error fetching news:', error));
  }



  const isLiked = async (postId: number) => {
    if (token) {
      try {
        const response = await fetch(`/api/news/${postId}/like/`, {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        const data = await response.json();
  
        return data.liked;
      } catch (error) {
        console.error('Error fetching like status:', error);
      }
    }
  
    return false;
  }

  const getDate = (created_at: number) => {
    const date = new Date(created_at * 1000);
    const months = ['янв', 'фев', 'мар', 'апр', 'май', 'июн', 'июл', 'авг', 'сен', 'окт', 'ноя', 'дек'];
    
    const day = date.getDate().toString().padStart(2, '0');
    const month = months[date.getMonth()];
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${day} ${month} ${hours}:${minutes}`;
  }

  const getHeshtags = (post: IPost) => {
    return JSON.parse(post.hashtags);
  }

  const handleLikeToPost = (postId: number) => {
    async function request(postId: number) {
      try {
        const response = await fetch(`/api/news/${postId}/like/`, {
          method: "POST",
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        const data = await response.json();

        if (data === "You have put a like") {
          return { liked: true };
        }
        else if (data === "You removed the like") {
          return { liked: false };
        }

      } catch (error) {
        console.error('Error fetching like status:', error);
      }

      return { liked: false }
    }

    if (token) {
      const responseData = request(postId);
      responseData.then(data => {
        setPosts([...posts.map(post => ({
          ...post,
          is_liked: data.liked,
          likes: post.id === postId ? (data.liked === true ? post.likes + 1 : post.likes - 1) : post.likes,
        }))])
      })
    }
    else {
      console.log("Not authorized")
    }
  }

  const getImage = (post: IPost) => {
    const images = JSON.parse(post.images);

    if (images.length > 0) {
      return images[0];
    }
    return "";
  }

  async function getEducation(userId: string) {
    async function request(userId: string) {
      try {
        const response = await fetch(`/api/universities/user/education/${userId}/`, {
          method: "GET"
        });
        const data = await response.json();
        
        return `${data.university.short_name}, ${data.department.short_name}, ${data.group.name}`
        
      } catch (error) {
        console.error('Error fetching like status:', error);
      }
  
      return "No University"
    }
    
    const education = await request(userId);
    
    return await education
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
                    <Link to={`/user/${post.author_id}`} style={{ textDecoration: "none", color: "black" }}>{post.author?.name} {post.author?.last_name} </Link>
                    <img src="/images/star.svg" alt="" /> <span>0</span>
                  </div>
                  <div className="date">{ getDate(post.created_at) }</div>
                </div>
                <div className="about__col col col-blue">
                  <span>{ post.author?.education }</span>
                </div>
              </div>
            </div>
            <div className="post__row row">
              { getImage(post) && <img src={getImage(post)} alt="" /> }
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
              <div>
                <button className="like" onClick={() => handleLikeToPost(post.id)}>
                  <img src="/images/like.svg" alt="" />
                </button>
                <span> {post.likes} </span>
              </div>
              <p className="blue-bold">{ post?.parsed_hashtags ? post?.parsed_hashtags[0] : "" }</p>
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
