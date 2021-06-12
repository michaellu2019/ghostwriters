import { useState, useEffect, useRef } from 'react';
import { useParams } from 'react-router-dom';

import Post from '../components/Post';
import FacebookLoginButton from '../components/FacebookLoginButton';

function Story(props) {
  const storyId = useParams().id;
  const [stories, setStories] = useState([]);
  const [errorMsg, setErrorMsg] = useState("");
  const postInput = useRef();
  
  useEffect(() => {
    getData();
  }, []);

  async function getData() {
    const requestOptions = {
      method: "GET",
      mode: "cors",
      headers: { "Content-Type": "application/json" }
    };

    const storiesResponse = await fetch("/api/get-story?id=" + storyId, requestOptions);
    const storiesData =  await storiesResponse.json();

    if (storiesData.status == "OK") {
      setStories([...storiesData.data.stories]);
    }
  }

  async function likePost(id) {
    if (props.loggedIn) {
      const postLike = {
        id: id,
        author: props.user.username + ":" + props.user.userId,
      };

      const requestOptions = {
        method: "POST",
        mode: "cors",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(postLike),
      };

      const response = await fetch(props.user.likedPosts.hasOwnProperty(id) ? "api/unlike-post" : "/api/like-post", requestOptions);
      const storiesData =  await response.json();

      if (storiesData.status == "OK") {
        getData();
        props.login(props.user);
        props.userLikePost(id);
      } else {
        setErrorMsg(storiesData.data);
      }
    }
  }

  async function submitPost(e) {
    e.preventDefault();

    const post = {
      story_id: parseInt(storyId),
      author: props.user.username + ":" + props.user.userId,
      text: postInput.current.value,
    };

    const requestOptions = {
      method: "POST",
      mode: "cors",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(post),
    };

    const response = await fetch("/api/create-post", requestOptions);
    const storiesData =  await response.json();

    if (storiesData.status == "OK") {
      getData();
      postInput.current.value = "";
      setErrorMsg("");
    } else {
      setErrorMsg(storiesData.data);
    }
  }

  return (
    <article className="story">
      {stories.map((story => 
        <div key={story.id}>
          <header>
            <h1><i>{story.title}</i></h1>
            <h3>by {story.author}</h3>
          </header>

          <div className="content">
            {story.content != null ? story.content.map((post => 
              <Post key={post.id} id={post.id} author={post.author} text={post.text} likes={post.likes} dislikes={post.dislikes} createdAt={post.created_at} liked={props.user.hasOwnProperty("likedPosts") && props.user.likedPosts.hasOwnProperty(post.id)} likePost={likePost} />
            )) : ""}
          </div>
        </div>
      ))}

      <div className="post-section">
        {props.loggedIn ? 
          <form>
            <span className="error-message">{errorMsg}</span>
            <textarea ref={postInput} placeholder={`Hi ${props.user.username}! Contribute to the story here!`}></textarea>
            <button onClick={submitPost}>SUBMIT</button>
          </form>
          :
          <div className="login-section">
            <h2>Want to contribute to the story?</h2>
            <FacebookLoginButton login={props.login} />
          </div>
        }
      </div>
    </article>
  );
}

export default Story;