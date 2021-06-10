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

    const response = await fetch("/api/get-story?id=" + storyId, requestOptions);
    const storiesData =  await response.json();

    if (storiesData.status == "OK") {
      setStories([...storiesData.data.stories]);
    }
  }

  async function submitPost(e) {
    e.preventDefault();

    const newPost = {
      story_id: parseInt(storyId),
      author: props.user.username,
      text: postInput.current.value,
    };

    console.log(newPost)

    const requestOptions = {
      method: "POST",
      mode: "cors",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(newPost),
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
        <div>
          <header>
            <h1><i>{story.title}</i></h1>
            <h3>by {story.author}</h3>
          </header>

          <div className="content">
            {story.content.map((post => 
              <Post key={post.id} id={post.id} author={post.author} text={post.text} likes={post.likes} dislikes={post.dislikes} createdAt={post.created_at}/>
            ))}
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


// import { useState, useEffect } from 'react';
// import { useParams } from 'react-router-dom';

// import Post from '../components/Post';

// function Story() {
//   const storyId = useParams().id;
//   const [story, setStory] = useState({
//     id: 0,
//     author: "",
//     title: "",
//     createdAt: "",
//     content: [],
//   });
//   const [storyContent, setStoryContent] = useState([]);
  
//   useEffect(() => {
//     getData();
//   }, []);

//   async function getData() {
//     const requestOptions = {
//       method: "GET",
//       mode: "cors",
//       headers: { "Content-Type": "application/json" }
//     };

//     const response = await fetch("/api/get-story?id=" + storyId, requestOptions);
//     const storyData =  await response.json();

//     if (storyData.status == "OK") {
//       // setStory(storyData.data.stories[0]);
//       setStoryContent([...storyContent, ...storyData.data.stories]);
//       console.log(storyData.data.stories)
//     }
//   }

//   return (
//     <article>
//       <header>
//         <h1><i>{story.title}</i></h1>
//         <h3>by {story.author}</h3>
//       </header>

//       <div className="content">
//         <div className="story-content">
//           {storyContent.map((post => {
//             <Post key={post.id} id={post.id} author={post.author} text={post.text} likes={story.likes} dislikes={story.dislikes} createdAt={story.createdAt}/>
//           }))}
//           {storyContent.map((post => {
//             <div>{post.author}</div>
//           }))}
//         </div>
//       </div>
//     </article>
//   );
// }

// export default Story;
