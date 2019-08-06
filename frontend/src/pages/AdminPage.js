import React, { Component } from 'react';
import '../css/Logging.css';
import { Button,ButtonGroup } from 'react-bootstrap';
import {getJwt} from "../jwt";

export class AdminPage extends Component {
  constructor(props) {
    super(props);
    this.handleEmail=this.handleEmail.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this)
    this.state = {
      email: '',
      password: ''
    }
  }

     handleSubmit (event) {
        event.preventDefault();
        console.log(getJwt());
        console.log(this.props.history.location)

        let registered;

        // perform all neccassary validations
        if (this.state.password === '') {
          alert("Please, fill in all fields!")
        } else {
          console.log('passed validation!')
          registered = true
          let data = JSON.stringify({
            Username: this.state.username,
            Password: this.state.password,
          });

          fetch('http://localhost:8080/login', {
            method: 'POST',
            body: data,
            headers:{
              'Content-Type': 'application/json'
            },
          }).then(res => res.json())
              .then(response => console.log('Body:', JSON.stringify(response)))
              .catch(error => console.error('Error:', error));;


          if (registered === true) {
            this.props.history.push('/logged');
          }

          console.log(getJwt());
        }

  }


  handleEmail(event) {
    this.setState({ username: event.target.value })
  }

  handlePasswordChange(event) {
    this.setState({ password: event.target.value })
  }

  render() {

    return (
        <div className="Registration">
          {/* main container */}
          <div className="reg">

            {/* Rectangle */}
            <div className="rectangle1">
              <div className="emptyRectangle1">
                <h2 className="rightTitle1">Admin Page</h2>
              </div>
            </div>
            {console.log(getJwt())}

            {/*Input form */}
            <div>
            <form className="reg-form"  onSubmit={this.handleSubmit}>
              <h4 className="font-weight-bold mb-3">Admin`s features</h4>
              <div>
                <h4>User</h4>

                <ButtonGroup className="mr-2 buttns" aria-label="First group">
                  <Button href="./usercreate" variant="secondary">Add user</Button>
                  <Button href="./userlist" variant="secondary">Get list</Button>
                  <Button href="./usermodify" variant="secondary">Modify user</Button>
                  <Button href="./userdelete" variant="secondary">Delete user</Button>
                </ButtonGroup>
                <h4>Post</h4>
                <ButtonGroup className="mr-2 buttns" aria-label="First group">
                  <Button href="./mainpage" variant="secondary">Add post</Button>
                  <Button href="./postlist" variant="secondary">Get list</Button>
                  <Button href="./postmodify" variant="secondary">Modify post</Button>
                  <Button href="./mainpage" variant="secondary">Delete post</Button>
                </ButtonGroup>
              </div>




              <h1 className="font-weight-bold greeting">Hello, Admin!</h1>
            </form>
            </div>


          </div>
          {/* end of main container */}
        </div>

    );
  }
}