import React  from 'react';
import '../css/Logging.css';
import { useEffect, useState } from 'react';
import {getJwt} from "../jwt";
import {Button, ButtonGroup} from "react-bootstrap";


export const PostList = () => {
  const [posts, setUsers] = useState([]);

  useEffect(() => {
    getUsers();
  }, []);

  async function getUsers() {
    let token = getJwt();
    console.log(token);
    const response = await fetch('http://localhost:8080/v1/posts',{
      method: 'GET',
      headers:{
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      }
    }).catch(error => console.error('Error:', error));
    const posts = await response.json();
    console.log(posts)
    setUsers(posts.posts)
  }

  return (
      <div className="Registration">
        {/* main container */}
        <div className="reg">

          {/* Rectangle */}
          <div className="rectangle1">
            <div className="emptyRectangle1">
              <h2 className="rightTitle1">Post List</h2>
            </div>
          </div>

          {/*Input form */}
          <div className="reg-form list-users" >
            <ul>
              {posts.map(post => (
                  <li>Title: {post.title}, Content: {post.content}, UserID: {post.id},
                  </li>
              ))}
            </ul>
          </div>

        </div>
        <ButtonGroup className="mr-2 buttns" aria-label="First group">
          <Button href="./mainpage" variant="secondary">Add post</Button>
          <Button href="./postlist" variant="secondary">Get list</Button>
          <Button href="./postmodify" variant="secondary">Modify post</Button>
          <Button href="./mainpage" variant="secondary">Delete post</Button>
        </ButtonGroup>
      </div>

  );
};

export default PostList;
