import React, { Component } from 'react';
import { getJwt } from '../jwt';
import axios from 'axios';
import { createBrowserHistory } from 'history';

export class Redirect extends Component {
    constructor(props) {
        super(props);

        this.state = {
            access: false
        }
    }

    componentDidMount() {

        const history = createBrowserHistory({ forceRefresh: true });
        const jwt = getJwt();

        if (!jwt) {

            setTimeout(function(){
                history.push('/')
            }.bind(this), 600);
            return;
        }

        axios.get('/getToken', { headers: { Authorization: `Bearer ${jwt}` } }).then(res => {

            this.setState({ access: true });
            return res;
        }).catch(err => {

            console.log(err.response);
            localStorage.removeItem('accessToken');
            localStorage.removeItem('refreshToken');

            setTimeout(function(){
                history.push('/')
            }.bind(this), 600);

        });
    }

    render() {
        if (this.state.access === false) {
            return (
                <div> </div>
            )
        }
        return (
            <div>
                {this.props.children}
            </div>
        )
    }
}