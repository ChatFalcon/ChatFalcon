import React from "react";
import saveConfig from "../saveConfig";

export default class UserConfiguration extends React.Component {
    componentDidMount() {
        this.props.setNextCb(this.nextButtonPressed.bind(this));
    }

    async nextButtonPressed() {
        this.props.config.stage = this.props.stage;
        await saveConfig(this.props.config);
    }

    render() {
        return <div>
            <h1 className="title is-1">User Configuration</h1>
            <p>
                Next we just simply need to create your first user. Please enter the information below.
            </p>
            <div className="field">
                <br />
                <label><b>E-mail Address:</b></label>
                <input className="input" type="email" placeholder="E-mail" value={this.props.config.firstUserEmail} onChange={e => this.props.config.firstUserEmail = e.target.value.trim()} />
            </div>
            <div className="field">
                <br />
                <label><b>Username:</b></label>
                <input className="input" type="text" placeholder="Username" value={this.props.config.firstUserUsername} onChange={e => this.props.config.firstUserUsername = e.target.value.trim()} />
            </div>
            <div className="field">
                <br />
                <label><b>Password:</b></label>
                <input className="input" type="password" placeholder="Password" value={this.props.config.firstUserPassword} onChange={e => this.props.config.firstUserPassword = e.target.value.trim()} />
            </div>
        </div>;
    }
}
