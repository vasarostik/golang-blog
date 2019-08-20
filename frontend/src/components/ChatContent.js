import React, { Component } from 'react';

class ChatContent extends Component {
  static createMarkup(markupString) {
    return {
      __html: markupString,
    }
  }

  render() {
    return (
      <div className="row">
        <div className="col s12">
          <div className="card horizontal">
            <div id="chat-messages" 
                 className="card-content" 
                 dangerouslySetInnerHTML={ChatContent.createMarkup(this.props.html)}>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default ChatContent; 
