import React, { Component } from 'react';
import emojione from 'emojione';
import ChatContent from '../components/ChatContent';
import '../css/Chat.css';
import ChatInput from '../components/ChatInput';
import {getJwt} from "../jwt";

export class Chat extends Component {
  constructor(props) {
    super(props);
    this.ws = null;
    this.state = {
      newMsg: '', // Holds new messages to be sent to the server
      chatContent: '', // A running list of chat messages displayed on the screen
      username: '', // Our username
    }
  }

  componentWillMount() {
    this.ws = new WebSocket('ws://localhost:8080/ws?token='+getJwt());
    this.ws.addEventListener('message', e => {
      let msg = JSON.parse(e.data);
      this.setState(prevState => {
        return {
          chatContent: prevState.chatContent + `<div class="chip color-green white-text"> ${msg.username} </div> ${emojione.toImage(msg.message)} <br/>`,
        }
      });
      const el = document.getElementById('chat-messages');
      el.scrollTop = el.scrollHeight; // auto scroll to bottom
    })
  }

  send() {
    if (this.state.newMsg !== '') {
      this.ws.send(
        JSON.stringify({
          email: this.state.email,
          username: this.state.username,
          message: this.state.newMsg, // strip out html
        })
      );
      // reset newMsg
      this.setState({
        newMsg: '',
      });
    }
  }

  updateMsg(e) {
    this.setState({
      newMsg: e.target.value,
    })
  }

  updateUsername(e) {
    this.setState({
      username: e.target.value,
    })
  }

  render() {

    return (
      <div >
        <header>
          <nav className="nav-wrapper color-green">
            <h2 className="logo center logotxt" >
              Chat
            </h2>
          </nav>
        </header>
        <main id="app">
          <ChatContent 
            html={this.state.chatContent}
          />
          <ChatInput
              value={this.state.newMsg}
              sendMessage={() => this.send()}
              updateMsg={e => this.updateMsg(e)}
              username={this.state.username}
              />
        </main>

      </div>
    );
  }
}

