import { useState, useEffect } from 'react';

import StoryCard from '../components/StoryCard';

function Bookshelf() {
  const [stories, setStories] = useState([]);
  
  useEffect(() => {
    getData();
  }, []);

  async function getData() {
    const requestOptions = {
      method: "GET",
      mode: "cors",
      headers: { "Content-Type": "application/json" }
    };

    const response = await fetch("/api/get-stories", requestOptions);
    const storiesData =  await response.json();

    if (storiesData.status == "OK") {
      setStories([...stories, ...storiesData.data.stories]);
      console.log(stories);
    }
  }

  return (
    <article>
      <header>
        <h1>Public Stories</h1>
      </header>

      <div className="content">
        <div className="stories-container">
          {stories.map((story => 
            <StoryCard key={story.id} id={story.id} author={story.author} title={story.title} content={story.content} createdAt={story.createdAt} />
          ))}
        </div>
      </div>
    </article>
  );
}

export default Bookshelf;
