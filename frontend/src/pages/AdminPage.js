import React, { Component } from 'react';
import '../css/Logging.css';
import { Button,ButtonGroup } from 'react-bootstrap';
import {getJwt} from "../jwt";
import {Chat} from "./Chat";

export class AdminPage extends Component {
  constructor(props) {
    super(props);
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
            <form className="reg-form">
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

          <ButtonGroup className="mr-2 buttns" aria-label="First group">
            <Button href="./chat" variant="secondary">Chat</Button>
          </ButtonGroup>

          {/* end of main container */}
        </div>

    );
  }
}