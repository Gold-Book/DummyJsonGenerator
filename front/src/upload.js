import React, { PureComponent } from "react";
import API from './api';

class Registory extends PureComponent {
  uploadData = event => {
      console.log(this.state)

      API.post("", {"heads" : this.state.headValue, "body" : this.state.headValue}).then(res => {
          console.log(res);
      });
  };

  handleInputHeader = event => {
      this.setState({headValue: event.target.value});
  }

  handleInputBody = event => {
      this.setState({bodyValue: event.target.value});
  }

  constructor(props) {
    super(props);
    this.state = {
      headValue: "",
      bodyValue: ""
    }
};

  render() {
    return (
        <div>
          <div className="level">
            <div className="level-left">
                <h1 className="level-item title">Upload Response Data</h1>
            </div>
            <div className="level-right">
                <span className="level-item button is-success is-selected" onClick={this.uploadData}>upload</span>
            </div>
          </div>
          <div className="field">
            <div className="control">
              <textarea className="textarea is-info" type="text" placeholder="input response header" value={this.state.headValue} onChange={this.handleInputHeader}></textarea>
            </div>
          </div>
          <div className="field">
            <div className="control">
              <textarea className="textarea is-info" type="text" placeholder="input response body" value={this.state.bodyValue} onChange={this.handleInputBody}></textarea>
            </div>
          </div>
      </div>
    );
  }
}

export default Registory;
