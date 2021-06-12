import { useState, useEffect } from 'react';
import { Route, Link } from 'react-router-dom';

import Bookshelf from './pages/Bookshelf';
import Story from './pages/Story';

import './styles/App.css';
import './styles/bookshelf.css';
import './styles/story.css';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [user, setUser] = useState({});

  async function login(user) {
    if (user.hasOwnProperty("userId") && user.hasOwnProperty("username")) {
      setLoggedIn(true);
      user.likedPosts = {};

      const requestOptions = {
        method: "GET",
        mode: "cors",
        headers: { "Content-Type": "application/json" }
      };

      const postLikesResponse = await fetch("/api/get-author-post-likes?author=" + user.username + ":" + user.userId, requestOptions);
      const postLikesData =  await postLikesResponse.json();
  
      console.log(postLikesData, "/api/get-author-post-likes?author=" + user.username + ":" + user.userId)
      if (postLikesData.status == "OK" && postLikesData.data.post_likes != null) {
        postLikesData.data.post_likes.forEach(post => {
          console.log(post.post_id, post.author)
          user.likedPosts[post.post_id] = post.author;
        });
      }
    }
    setUser(user);
  }

  function userLikePost(id) {
    if (user.hasOwnProperty("likedPosts")) {
      setUser((oldUser => ({
        ...oldUser,
        [oldUser.likedPosts]: {
          ...oldUser.likedPosts, 
          [id]: oldUser.username + ":" + oldUser.userId,
        },
      })));
    }
  }

  return (
    <div className="wrapper">
      <header className="primary">
          <nav>
            <span className="heading"><Link className="nav-link" to="/">Ghostwriters</Link></span>

            <ul className="nav-buttons links">
              <li><Link className="nav-link" to="/">Stories</Link></li>
            </ul>
          </nav>
        </header>

      <main className="primary">
        <Route exact path="/" render={() => <Bookshelf loggedIn={loggedIn} user={user} login={login} userLikePost={userLikePost} />} />
        <Route exact path="/story/:id" render={() => <Story loggedIn={loggedIn} user={user} login={login} userLikePost={userLikePost} />} />
      </main>
    </div>
  );
}

export default App;
