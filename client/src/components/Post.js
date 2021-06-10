import { useState } from 'react';

function Post(props) {
  const [active, setActive] = useState(false);

  const postTimestamps = props.createdAt.split(/[-TZ:]/);
  const postDate = new Date(Date.UTC(postTimestamps[0], postTimestamps[1], postTimestamps[2], postTimestamps[3], postTimestamps[4], postTimestamps[5]));
  const formattedPostDate = `${(postDate.getMonth())}/${+ postDate.getDate()}/${postDate.getFullYear()}`;

  return (
    <div className="post">
      <div className="post-text" onMouseEnter={() => setActive(true)} onMouseLeave={() => setActive(false)}>{props.text}</div>
      <div className={active ? "post-info active" : "post-info"}>
        <div>Written by {props.author} on {formattedPostDate + " - " + postDate.toLocaleTimeString()}</div>
      </div>
    </div>
  );
}

export default Post;
