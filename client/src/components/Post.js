import { useState } from 'react';

function Post(props) {
  const [active, setActive] = useState(false);

  const postTimestamps = props.createdAt.split(/[-TZ:]/);
  const postDate = new Date(Date.UTC(postTimestamps[0], postTimestamps[1], postTimestamps[2], postTimestamps[3], postTimestamps[4], postTimestamps[5]));
  const formattedPostDate = `${(postDate.getMonth())}/${+ postDate.getDate()}/${postDate.getFullYear()}`;

  let author = props.author;
  if (author.indexOf(":") > -1) {
    author = author.split(":")[0];
  }

  return (
    <div className="post">
      <div className="post-text" onMouseEnter={() => setActive(true)} onMouseLeave={() => setActive(false)}>
        <span className="heart-container">
          <svg className={props.liked ? "heart-button active" : "heart-button"} onClick={() => props.likePost(props.id)}><path d = "M17.027 2.21c-2.248 0-4.166 1.786-5.027 3.704C11.139 3.995 9.222 2.21 6.973 2.21 3.931 2.21 1.416 4.725 1.416 7.766c0 6.218 6.283 7.872 10.584 14.024 4.035-6.152 10.584-8.005 10.584-14.024C22.584 4.725 20.072 2.21 17.027 2.21z" /></svg>
          {props.likes}
        </span>
        <span>{props.text}</span>
      </div>
      <div className={active ? "post-info active" : "post-info"}>
        <div>Written by {author} on {formattedPostDate + " - " + postDate.toLocaleTimeString()}</div>
      </div>
    </div>
  );
}

export default Post;
