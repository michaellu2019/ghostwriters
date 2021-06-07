import { useState, useEffect } from 'react';

function Post(props) {
  return (
    <div className="post">
      <p>{props.text}</p>
      <span>{props.author} - {props.likes}, {props.dislikes}</span>
    </div>
  );
}

export default Post;
