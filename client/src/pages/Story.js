import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

import Post from '../components/Post';

function Story() {
  const storyId = useParams().id;
  const [story, setStory] = useState({
    id: 0,
    author: "",
    title: "",
    createdAt: "",
    content: [],
  });
  const [storyContent, setStoryContent] = useState([]);
  
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
    const storyData =  await response.json();

    if (storyData.status == "OK") {
      setStory(storyData.data.stories[0]);
      setStoryContent([...storyContent, ...storyData.data.stories[0].content]);
    }
  }

  return (
    <article>
      <header>
        <h1><i>{story.title}</i></h1>
        <h3>by {story.author}</h3>
      </header>

      <div className="content">
        <div className="story-content">
          {storyContent.map((post => {
            <Post key={post.id} id={post.id} author={post.author} text={post.text} likes={story.likes} dislikes={story.dislikes} createdAt={story.createdAt}/>
          }))}
        </div>
      </div>
    </article>
  );
}

export default Story;
