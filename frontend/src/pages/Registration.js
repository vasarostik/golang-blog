import React, { Component } from 'react';
import { MDBInput } from 'mdbreact';
import '../css/Logging.css';
import "babel-polyfill";
import {getJwt, getRefreshToken} from '../jwt';
import {Button, ButtonGroup} from "react-bootstrap";

export class Registration extends Component {
    constructor(props) {
        super(props);
        this.handleFirstNameChange=this.handleFirstNameChange.bind(this);
        this.handleLastNameChange=this.handleLastNameChange.bind(this);
        this.handlePasswordConfirmChange=this.handlePasswordConfirmChange.bind(this);
        this.handleUsername=this.handleUsername.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleLogin = this.handleLogin.bind(this);
        this.state = {
            firstName:'',
            lastName:'',
            userName:'',
            password:'',
            passwordConfirm:'',
            RoleID: 200
        }
    }


    async handleLogin () {

        let registered;

        // perform all neccassary validations
        if (this.state.password === '') {
            alert("Please, fill in all fields!")
        } else {
            console.log('passed validation!')
            registered = true
            let data = JSON.stringify({
                Username: this.state.userName,
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

    async handleSubmit (event) {
        event.preventDefault();
        let token = getJwt();
        console.log(token);
        console.log(this.props.history.location);

        let registered;

        // perform all neccassary validations
        if (this.state.password === '') {
            alert("Please, fill in all fields!")
        } else if (this.state.password !== this.state.passwordConfirm) {
            alert("Passwords don't match");

        } else {
            console.log('passed validation!');
            registered = true;
            let data = JSON.stringify({
                first_name: this.state.firstName,
                last_name: this.state.lastName,
                Username: this.state.userName,
                role_id: this.state.RoleID,
                Password: this.state.password,
                password_confirm: this.state.passwordConfirm
            });

            const response = await fetch('http://localhost:8080/users', {
                method: 'POST',
                body: data,
                headers:{
                    'Content-Type': 'application/json',
                }
            }).catch(error => console.error('Error:', error));


            if (response !== undefined) {

                console.log('Status: ' + response.status);

                if (response.status === 200) {
                    registered = true;

                    const json = await response.json();

                    console.log('success');
                } else {
                    registered = false;
                }
            }

            if (registered === true) {
                this.handleLogin();
            }
        }

    }

    handleFirstNameChange(event) {
        this.setState({ firstName: event.target.value })
    }

    handleLastNameChange(event) {
        this.setState({ lastName: event.target.value })
    }

    handleUsername(event) {
        this.setState({ userName: event.target.value })
    }

    handlePasswordChange(event) {
        this.setState({ password: event.target.value })
    }

    handlePasswordConfirmChange(event) {
        this.setState({ passwordConfirm: event.target.value })
    }


    render() {

        return (
            <div className="Registration">
                {/* main container */}
                <div className="reg">

                    {/* Rectangle */}
                    <div className="rectangle1">
                        <div className="emptyRectangle1">
                            <h2 className="rightTitle1">Create account</h2>
                        </div>
                    </div>


                    {/*Input form */}
                    <form className="reg-form"  onSubmit={this.handleSubmit}>
                        <h4 className="font-weight-bold mb-3">Registration form</h4>

                        <div className="reg-input">
                            <MDBInput label="First Name" outline icon="envelope" onChange={this.handleFirstNameChange.bind(this)} required />
                        </div>

                        <div className="reg-input">
                            <MDBInput label="Last Name" outline icon="envelope" onChange={this.handleLastNameChange.bind(this)} required />
                        </div>

                        <div className="reg-input">
                            <MDBInput label="Username" outline icon="envelope" onChange={this.handleUsername.bind(this)} required />
                        </div>

                        <div className="reg-input">
                            <MDBInput label="Password" type="password" outline icon="fas fa-key" onChange={this.handlePasswordChange.bind(this)} required />
                        </div>

                        <div className="reg-input">
                            <MDBInput label="Confirm Password" type="password" outline icon="fas fa-key" onChange={this.handlePasswordConfirmChange.bind(this)} required />
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
                    <Button href="./" variant="secondary">Log in</Button>
                </ButtonGroup>

            </div>

        );
    }
}