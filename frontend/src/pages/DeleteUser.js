import React, {Component} from 'react';
import { MDBInput } from 'mdbreact';
import '../css/Logging.css';
import {Button, ButtonGroup} from "react-bootstrap";
import {getJwt} from "../jwt";

export class DeleteUser extends Component {
  constructor(props) {
    super(props);
    this.handleID=this.handleID.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.state = {
      ID: ''
    }
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
          });
          console.log(data);
          const response = await fetch('http://localhost:8080/v1/users/'+this.state.ID, {
            method: 'DELETE',
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

  render() {

    return (
        <div className="Registration">
          {/* main container */}
          <div className="reg">

            {/* Rectangle */}
            <div className="rectangle1">
              <div className="emptyRectangle1">
                <h2 className="rightTitle1">Delete user</h2>
              </div>
            </div>

            {/*Input form */}
            <form className="reg-form"  onSubmit={this.handleSubmit}>
              <h4 className="font-weight-bold mb-3">Delete User</h4>
              <div className="reg-input">
                <MDBInput label="ID" outline icon="envelope" onChange={this.handleID.bind(this)} required />
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
          <ButtonGroup className="mr-2 buttns" aria-label="First group">
            <Button href="./adminpage" variant="secondary">Back</Button>
          </ButtonGroup>
          {/* end of main container */}
        </div>

    );
  }
}