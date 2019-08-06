import React, {Component} from 'react';
import { MDBInput } from 'mdbreact';
import '../css/Logging.css';
import {getJwt} from "../jwt";
import {Button, ButtonGroup} from "react-bootstrap";

export class ModifyUser extends Component {
  constructor(props) {
    super(props);
    this.handleID=this.handleID.bind(this);
    this.handleFirstName=this.handleFirstName.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleLastName = this.handleLastName.bind(this);
    this.getUsers1 = this.getUsers1.bind(this);
    this.state = {
      ID: '',
      FirstName: '',
      LastName: '',
      users: []
    }
  }

  async getUsers1() {
    let token = getJwt();
    console.log(token);
    const response = await fetch('http://localhost:8080/v1/users/'+this.state.ID,{
      method: 'GET',
      headers:{
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`
      }
    }).catch(error => console.error('Error:', error));
    const users1 = await response.json();
    console.log(users1);
    this.setState(this.state.users[0]=users1)
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
            first_name: this.state.FirstName,
            last_name: this.state.LastName,
          });
          console.log(data);
          const response = await fetch('http://localhost:8080/v1/users/'+this.state.ID, {
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
              this.getUsers1();

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

  handleFirstName(event) {
    this.setState({ FirstName: event.target.value })
  }

  handleLastName(event) {
    this.setState({ LastName: event.target.value })
  }



  render() {



    return (
        <div className="Registration">
          {/* main container */}
          <div className="reg">

            {/* Rectangle */}
            <div className="rectangle1">
              <div className="emptyRectangle1">
                <h2 className="rightTitle1">Modify user</h2>
              </div>
            </div>

            {/*Input form */}
            <form className="reg-form"  onSubmit={this.handleSubmit}>
              <h4 className="font-weight-bold mb-3">Modify User</h4>
              <div className="reg-input">
                <MDBInput label="ID" outline icon="envelope" onChange={this.handleID.bind(this)} required />
              </div>
              <div className="reg-input">
                <MDBInput label="First Name" outline icon="envelope" onChange={this.handleFirstName.bind(this)} required />
              </div>
              <div className="reg-input">
                <MDBInput label="Last Name" type="text" outline icon="fas fa-key" onChange={this.handleLastName.bind(this)} required />
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
                {this.state.users.map(user => (
                    <li>ID: {user.id}, First Name: {user.first_name}, Last Name: {user.last_name}</li>
                ))}
              </ul>
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