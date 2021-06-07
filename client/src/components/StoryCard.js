import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

function StoryCard(props) {
  return (
    <Link to={"/story/" + props.id}>
      <div className="story-card">
        <header>
          <h3><i>{props.title}</i></h3>
          <h5>by {props.author}</h5>
        </header>
        <img src="" />
        <footer>Posts: {props.content.length}</footer>
      </div>
    </Link>
  );
}

export default StoryCard;
