import React, {Component, useEffect, useState} from 'react';
import { MDBInput } from 'mdbreact';
import '../css/Logging.css';
import {getJwt} from "../jwt";
import {Button, ButtonGroup} from "react-bootstrap";


export const PostMyList = () => {
  const [posts, setUsers] = useState([]);

  useEffect(() => {
    getUsers();
  }, []);

  async function getUsers() {
    let token = getJwt();
    console.log(token);
    const response = await fetch('http://localhost:8080/v1/posts/my',{
      method: 'GET',
      headers:{
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      }
    }).catch(error => console.error('Error:', error));
    const posts = await response.json();
    console.log(posts);
    setUsers(posts.posts)
  }

  return (
      <div>


          {/*Input form */}
          <div className="reg-form list-posts" >
            <ul>
              {posts.map(user => (
                  <li>Post ID: {user.post_id}, Title: {user.title}, Content: {user.content}</li>
              ))}
            </ul>

        </div>

      </div>

  );
};


export class ModifyPost extends Component {
  constructor(props) {
    super(props);
    this.handleID=this.handleID.bind(this);
    this.handleContent=this.handleContent.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleTitle = this.handleTitle.bind(this);
    this.getPosts1 = this.getPosts1.bind(this);
    this.state = {
      ID: '',
      title: '',
      content: '',
      posts: [],
    };
  }

  async getPosts1() {
    let token = getJwt();
    console.log(token);
    const response = await fetch('http://localhost:8080/v1/post/'+this.state.ID,{
      method: 'GET',
      headers:{
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      }
    }).catch(error => console.error('Error:', error));
    const posts1 = await response.json();
    console.log(posts1);
    this.setState(this.state.posts[0]=posts1)
  }

     async handleSubmit (event) {
        event.preventDefault();
        let token = getJwt();
        console.log(token);
        console.log(this.props.history.location);

        let registered;

        // perform all neccassary validations
        if (this.state.ID === '') {
          alert("Please, fill in all fields!")
        } else {
          console.log('passed validation!');
          registered = true;
          let data = JSON.stringify({
            ID: this.state.ID,
            title: this.state.title,
            content: this.state.content,
          });
          console.log(data);
          const response = await fetch('http://localhost:8080/v1/post/'+this.state.ID, {
            method: 'PATCH',
            body: data,
            headers:{
              'Content-Type': 'application/json',
              mode: 'CORS',
              Authorization: `Bearer ${token}`

            }
          }).catch(error => console.error('Error:', error));


          if (response !== undefined) {

            console.log('Status: ' + response.status);

            if (response.status === 200) {

              const json = await response.json();

              console.log(getJwt());
              console.log(json);
              this.getPosts1();

              console.log('success');
            } else {
              alert('error');
            }
          }
        }
   }




  handleID(event) {
    this.setState({ ID: event.target.value })
  }

  handleTitle(event) {
    this.setState({ title: event.target.value })
  }

  handleContent(event) {
    this.setState({ content: event.target.value })
  }



  render() {
    return (
        <div className="Registration">
          {/* main container */}
          <div className="reg">

            {/* Rectangle */}
            <div className="rectangle1">
              <div className="emptyRectangle1">
                <h2 className="rightTitle1">Modify post</h2>
              </div>
            </div>

            {/*Input form */}
            <form className="reg-form"  onSubmit={this.handleSubmit}>
              <h4 className="font-weight-bold mb-3">Modify post form</h4>
              <div className="reg-input">
                <MDBInput label="ID" outline icon="envelope" onChange={this.handleID.bind(this)} required />
              </div>
              <div className="reg-input">
                <MDBInput label="Title" outline icon="envelope" onChange={this.handleTitle.bind(this)} required />
              </div>
              <div className="reg-input">
                <MDBInput label="Content" type="text" outline icon="fas fa-key" onChange={this.handleContent.bind(this)} required />
              </div>


              <div className="space">
                <div className="float-left">
                  <a href="/" ><button className="signup-but cancel" type="button" style={{ color: 'white' }}>Cancel</button></a>
                </div>
                <div className="float-right">
                  <button className="signup-but sign-up" type="submit" style={{ color: 'white' }}>Continue</button>
                </div>
              </div>
              <ul>
                {this.state.posts.map(post => (
                    <li>Post ID: {post.post_id}, Title: {post.title}, Content: {post.content}</li>
                ))}
              </ul>
            </form>


          </div>
          <PostMyList />
          <ButtonGroup className="mr-2 buttns" aria-label="First group">
            <Button href="./mainpage" variant="secondary">Back</Button>
          </ButtonGroup>
          {/* end of main container */}
        </div>

    );
  }
}