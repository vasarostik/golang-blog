import React, { Component } from 'react';
import { MDBInput } from 'mdbreact';
import '../css/Logging.css';
import ListGroup from 'react-bootstrap/ListGroup';
import "babel-polyfill";
import { getJwt,getRefreshToken } from '../jwt';
import {Button, ButtonGroup} from "react-bootstrap";

export class Logging extends Component {
  constructor(props) {
    super(props);
    this.handleUsername=this.handleUsername.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this)
    this.state = {
      email: '',
      password: ''
    }
  }

    async handleSubmit (event) {
        event.preventDefault();

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

          const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            body: data,
            headers:{
              'Content-Type': 'application/json'
            }
          }).catch(error => console.error('Error:', error));


            if (response !== undefined) {

              console.log('Status: ' + response.status);

              if (response.status === 200) {
                registered = true;

                const json = await response.json();
                await localStorage.setItem('accessToken', json.token);
                await localStorage.setItem('refreshToken', json.refresh_token);

                console.log(getJwt());
                console.log(getRefreshToken());

                console.log('success');
              } else {
                registered = false;
                alert('Wrong username or password');
              }
            }



          if (registered === true) {
            this.state.username === 'admin' ? this.props.history.push('/adminpage') :
            this.props.history.push('/mainpage');
          }
        }

  }


  handleUsername(event) {
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
                <h2 className="rightTitle1">Log in Form</h2>
              </div>
            </div>


            {/*Input form */}
            <form className="reg-form"  onSubmit={this.handleSubmit}>
              <h4 className="font-weight-bold mb-3">Log in</h4>
              <p className="mdb-color-text">To sign in, please fill in these text fields</p>

              <div className="reg-input">
                <MDBInput label="Username" outline icon="envelope" onChange={this.handleUsername.bind(this)} required />
              </div>
              <div className="reg-input">
                <MDBInput label="Password" type="password" outline icon="fas fa-key" onChange={this.handlePasswordChange.bind(this)} required />
              </div>


              <div className="space">
                <div className="float-left">
                  <a href="/" ><button className="signup-but cancel" type="button" style={{ color: 'white' }}>Cancel</button></a>
                </div>
                <div className="float-right">
                  <button className="signup-but sign-up" type="submit" style={{ color: 'white' }}>Continue</button>
                </div>
              </div>
              <ButtonGroup className="mr-2 buttns" aria-label="First group">
                <Button href="./registration" variant="secondary">Sign up</Button>
              </ButtonGroup>
            </form>

          </div>
          {/* end of main container */}
        </div>

    );
  }
}