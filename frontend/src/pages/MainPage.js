import React, { Component } from 'react';
import { MDBInput } from 'mdbreact';
import '../css/Logging.css';
import {getJwt} from "../jwt";
import {Button, ButtonGroup} from "react-bootstrap";

export class MainPage extends Component {
  constructor(props) {
    super(props);
    this.handleEmail=this.handleEmail.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this)
    this.state = {
      title: '',
      content: ''
    }
  }

     handleSubmit (event) {
        event.preventDefault();
        let token = getJwt();
        console.log(token);
        console.log(this.props.history.location);

        let registered;


          console.log('passed validation!');
          registered = true;
          let data = JSON.stringify({
            title: this.state.title,
            content: this.state.content,
          });

          fetch('http://localhost:8080/v1/post/create', {
            method: 'POST',
            body: data,
            headers:{
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`

            }
          }).then(res => res.json())
              .then(response => console.log('Body:', JSON.stringify(response)))
              .catch(error => console.error('Error:', error));;


          if (registered === true) {
            this.props.history.push('/postlist');
          }

          console.log(getJwt());


  }


  handleEmail(event) {
    this.setState({ title: event.target.value })
  }

  handlePasswordChange(event) {
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
                <h2 className="rightTitle1">Main Page</h2>
              </div>
            </div>



            {/*Input form */}
            <form className="reg-form"  onSubmit={this.handleSubmit}>
              <h4 className="font-weight-bold mb-3">Add new Post</h4>

              <div className="reg-input">
                <MDBInput label="Title" outline icon="envelope" onChange={this.handleEmail.bind(this)} required />
              </div>
              <div className="reg-input">
                <MDBInput label="Content" type="textarea" outline icon="fas fa-key" onChange={this.handlePasswordChange.bind(this)} required />
              </div>


              <div className="space">
                <div className="float-left">
                  <a href="/" ><button className="signup-but cancel" type="button" style={{ color: 'white' }}>Cancel</button></a>
                </div>
                <div className="float-right">
                  <button className="signup-but sign-up" type="submit" style={{ color: 'white' }}>Continue</button>
                </div>
              </div>

            </form>


          </div>
          {/* end of main container */}
          <ButtonGroup className="mr-2 buttns" aria-label="First group">
            <Button href="./mainpage" variant="secondary">Add post</Button>
            <Button href="./postlist" variant="secondary">Get list</Button>
            <Button href="./postmodify" variant="secondary">Modify post</Button>
            <Button href="./mainpage" variant="secondary">Delete post</Button>
          </ButtonGroup>

          <ButtonGroup className="mr-2 buttns" aria-label="First group">
            <Button href="./chat" variant="secondary">Chat</Button>
          </ButtonGroup>
        </div>

    );
  }
}